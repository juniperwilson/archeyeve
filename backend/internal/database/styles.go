package database

import (
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
)

type Style struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AddStyle(name string) {

	var id int
	err := dbpool.QueryRow(context.Background(), `INSERT INTO styles (name) VALUES ($1) RETURNING id`, name).Scan(&id)
	if err != nil {
		panic("failed to insert style: " + err.Error())
	}
	_, err = dbpool.Exec(context.Background(), `INSERT INTO paths (style_id, path) VALUES ($1, $2)`, id, fmt.Sprint(id))
	if err != nil {
		panic("failed to insert style path: " + err.Error())
	}
}

func AddEdge(parent, child int) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select()
	_, err := dbpool.Exec(context.Background(), `SELECT add_edge($1, $2)`, parent, child)
	if err != nil {
		panic("failed to add edge: " + err.Error())
	}
}

func DeleteEdge(parent, child int) {
	_, err := dbpool.Exec(context.Background(), `SELECT delete_edge($1, $2)`, parent, child)
	if err != nil {
		panic("failed to delete edge: " + err.Error())
	}
}

func FindStyle(name string) (Style, error) {
	var id int
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id").From("styles")

	if name != "" {
		sb.Where(sb.EQ("name", name))
	}
	sql, args := sb.Build()
	println(sql)

	err := dbpool.QueryRow(context.Background(), sql, args...).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return Style{}, err
	}

	return Style{id, name}, nil
}

func GetAllStyles() ([]*Style, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(`*`).From(`styles`)

	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	fmt.Println(sql)
	fmt.Println(args)
	rows, err := dbpool.Query(context.Background(), sql, args...)
	if err != nil {
		fmt.Println(err.Error())
	}

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (*Style, error) {
		var s Style
		err := row.Scan(&s.ID, &s.Name)
		return &s, err
	})
}
