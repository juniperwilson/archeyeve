package database

import (
	"context"

	"github.com/huandu/go-sqlbuilder"
)

type Observation struct {
	ID       int      `json:"id"`
	UserID   int      `json:"userid"`
	Name     string   `json:"name"`
	Lng      float64  `json:"lng"`
	Lat      float64  `json:"lat"`
	Styles   []string `json:"styles"`
	Year     int      `json:"year"`
	ImgCount int      `json:"imgcount"`
}

func AddObservation(obs *Observation) (int, error) {

	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	ib.InsertInto("observations")
	ib.Cols("userid", "name", "lng", "lat", "style_ids", "year", "img_count")
	ib.Values(obs.UserID, obs.Name, obs.Lng, obs.Lat, obs.Styles, obs.Year, obs.ImgCount)
	ib.Returning("id")

	sql, args := ib.Build()
	id := 0
	err := dbpool.QueryRow(context.Background(), sql, args).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
