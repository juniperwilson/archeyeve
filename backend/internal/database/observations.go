package database

import (
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

type Observation struct {
	ID       int      `json:"id" db:"id" fieldtag:"id"`
	UserID   int      `json:"userid" db:"userid"`
	Name     string   `json:"name" db:"name"`
	Lng      float64  `json:"lng" db:"lng"`
	Lat      float64  `json:"lat" db:"lat"`
	Styles   []string `json:"styles" db:"style_ids"`
	Year     int      `json:"year" db:"year"`
	ImgCount int      `json:"imgcount" db:"imgcount"`
}

func AddObservation(obs *Observation) (int, error) {

	dbObservation := sqlbuilder.NewStruct(obs)

	ib := dbObservation.WithoutTag("id").InsertInto("observations", obs)
	sql, args := ib.BuildWithFlavor(sqlbuilder.PostgreSQL)

	id := 0
	fmt.Println(sql)
	fmt.Println(args)
	_, err := dbpool.Exec(context.Background(), sql, args...)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	fmt.Println(id)
	return id, nil
}
