package database

import (
	"context"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
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

func FindStyle(name string) Style {
	var id int
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("id").From("styles").Where(sb.EQ("name", name))
	sql, args := sb.Build()

	err := dbpool.QueryRow(context.Background(), sql, args...).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to find style: " + err.Error())
	}

	return Style{id, name}
}
