package main

import (
	_ "embed"
	"fmt"
	"strings"

	"slices"

	"github.com/juniperwilson/archeyeve/external/api"
	"github.com/juniperwilson/archeyeve/internal/database"
)

//go:embed styles.txt
var Styles string

var loneStyles = []string{"megalithic", "roman"}

func SeedAll() {
	seedStyles()
	seedObservations()
}

func seedStyles() {
	styles := strings.Split(Styles, "\n")
	s := [][]string{}
	for _, v := range styles {
		if v != "" {
			s = append(s, strings.Split(v, " --> "))
		}
	}
	flatS := []string{}

	for _, v := range s {
		if !slices.Contains(flatS, v[0]) {
			flatS = append(flatS, v[0])
		}
		if !slices.Contains(flatS, v[1]) {
			flatS = append(flatS, v[1])
		}
	}

	for _, v := range flatS {
		database.AddStyle(strings.ToLower(v))
	}

	seedEdges(s, flatS)
}

func seedEdges(s [][]string, flatS []string) {
	for _, v := range s {
		index1 := slices.Index(flatS, v[0])
		index2 := slices.Index(flatS, v[1])

		database.AddEdge(index1+1, index2+1)
	}
}

func seedObservations() {

	obs := []database.Observation{{
		ID:       0,
		UserID:   0,
		Name:     "Praça de Touros do Campo Pequeno",
		Lng:      -9.145202,
		Lat:      38.742590,
		Usage:
		Styles:   []string{"neo-mudejar"},
		Year:     1892,
		ImgCount: 1,
	}, {
		ID:       1,
		UserID:   0,
		Name:     "Torre do Tombo",
		Lng:      -9.156460,
		Lat:      38.754630,
		Styles:   []string{"post-modernism", "brutalism"},
		Year:     1990,
		ImgCount: 1,
	}, {
		ID:       2,
		UserID:   0,
		Name:     "Torre da Praça do Areeiro",
		Lng:      -9.133270,
		Lat:      38.742910,
		Styles:   []string{"português suave"},
		Year:     1938,
		ImgCount: 1,
	}, {
		ID:       3,
		UserID:   0,
		Name:     "St George's Church",
		Lng:      -9.160482,
		Lat:      38.716496,
		Styles:   []string{"neo-romanesque"},
		Year:     1889,
		ImgCount: 1,
	}, {
		ID:       4,
		UserID:   0,
		Name:     "EDP Sede II",
		Lng:      -9.148990,
		Lat:      38.707162,
		Styles:   []string{"brutalism"},
		Year:     2024,
		ImgCount: 1,
	}, {
		ID:       5,
		UserID:   0,
		Name:     "Igreja de São João Baptista",
		Lng:      -8.414697,
		Lat:      39.603648,
		Styles:   []string{"manueline"},
		Year:     1500,
		ImgCount: 1,
	}, {
		ID:       6,
		UserID:   0,
		Name:     "Sé de Braga",
		Lng:      -8.427396,
		Lat:      41.549916,
		Styles:   []string{"romanesque", "manueline", "baroque"},
		Year:     1089,
		ImgCount: 1,
	}, {
		ID:       7,
		UserID:   0,
		Name:     "Santuário do Monte de Santa Luzia",
		Lng:      -8.835104,
		Lat:      41.701517,
		Styles:   []string{"neo-byzantine", "neo-romanesque", "neo-gothic"},
		Year:     1904,
		ImgCount: 1,
	}, {
		ID:       8,
		UserID:   0,
		Name:     "Citânia de Briteiros",
		Lng:      -8.315402,
		Lat:      41.527687,
		Styles:   []string{"castro", "iron-age"},
		Year:     -200,
		ImgCount: 1,
	}}

	for _, v := range obs {
		err := api.SeedObservation(&v)
		if err != nil {
			fmt.Println(err.Error())
		}

	}
}
