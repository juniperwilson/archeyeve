package database

import (
	"context"
	"fmt"
	"math"
	"github.com/jackc/pgx/v5"

	"github.com/huandu/go-sqlbuilder"
)
const BUILDING_RADIUS = 0.001
const ROUNDING_LIMIT = 0.000499
type Observation struct {
	ID       int      `json:"id" db:"id" fieldtag:"id"`
	UserID   int      `json:"userid" db:"userid"`
	Name     string   `json:"name" db:"name"`
	Address  string   `json:"address" db:"address"`
	Type	 []string
	Lng      float64  `json:"lng" db:"lng"`
	Lat      float64  `json:"lat" db:"lat"`
	Styles   []string `json:"styles" db:"styles"`
	Year     int      `json:"year" db:"year"`
	ImgCount int      `json:"imgcount" db:"imgcount"`
}

func AddObservation(obs *Observation) (int, error) {

	dbObservation := sqlbuilder.NewStruct(obs)

	ib := dbObservation.WithoutTag("id").InsertInto("observations", obs)
	sql, args := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)

	id := 0
	_, err := dbpool.Exec(context.Background(), sql, args...)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return id, nil
}

func GetObservation(id int) (Observation, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("*").From("observations").Where(sb.EQ("id", id))

	obs := Observation{}
	sql, args := sb.Build()
	err := dbpool.QueryRow(context.Background(), sql, args...).Scan(&obs.ID, &obs.UserID, &obs.Name, &obs.Address, &obs.Lng, &obs.Lat, &obs.Styles, &obs.Year, &obs.ImgCount)
	if err != nil {
		return Observation{}, err
	}

	return obs, nil
}

//returns list of observations with building tag within a certain radius
// of the coords given
func GetBuildings(lat, lng float64, name string) ([]*Observation, error) {
	roundLat := math.Round(lat * 1000) / 1000

	roundLng := math.Round(lng * 1000) / 1000
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("*").From("observations").Where(sb.EQ("isBuilding", "true"))
	sb.Between("lat", roundLat, roundLat + ROUNDING_LIMIT)
	sb.Between("lng", roundLng, roundLng + ROUNDING_LIMIT)
	
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	fmt.Println(sql)
	fmt.Println(args)
	rows, err := dbpool.Query(context.Background(), sql, args...)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(rows)

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (*Observation, error) {
		var o Observation
		err := row.Scan(&o.ID, &o.UserID, &o.Name, &o.Address, &o.Lng, &o.Lat, &o.Styles, &o.Year, &o.ImgCount)
		fmt.Println(o)
		return &o, err
	})

}
