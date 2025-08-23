package main

import (
	"github.com/juniperwilson/archeyeve/internal/database"
)

func SeedAll() {
	seedStyles()
}

func seedStyles() {
	values := []string{
		"Portuguese",
		"Gothic",
		"Mudejar",
		"Manueline",
		"Neo-Manueline",
	}

	edges := [][2]int{
		{1, 2},
		{1, 3},
		{2, 4},
		{3, 4},
		{4, 5},
	}

	for _, v := range values {
		database.AddStyle(v)
	}

	for _, e := range edges {
		database.AddEdge(e[0], e[1])
	}
}
func seedObservations() {
	
}
