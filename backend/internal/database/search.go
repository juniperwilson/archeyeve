package database

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
)

type Search struct {
	Style string `json:"style"`
	//Coord 1 represents NW corner of selection, or center
	Lng1 float64 `json:"lng1"`
	Lat1 float64 `json:"lat1"`
	//Coord 2 represents SE corner of selection
	Lng2    float64 `json:"lng2"`
	Lat2    float64 `json:"lat2"`
	Radius  float64 `json:"radius"`
	BefYear int     `json:"befyear"`
	AftYear int     `json:"aftyear"`
}

// POSSIBLE SEARCH TYPES
// Square search, both coords, no radius
// Round search, one coord with radius
// Style is optional
// years are optional
func Find(s *Search) ([]*Observation, error) {

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(`*`).From(`observations`)

	style, err := styleSearch(sb, s)
	if err != nil {
		return nil, err
	}
	era := temporalSearch(sb, s)
	region := spatialSearch(sb, s)

	//Valid search check
	if region == "any" && style == "any" && era == "any" {
		sb = sqlbuilder.NewSelectBuilder()
		sb.Select("*").From("observations")
	}

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
		err := row.Scan(&o.ID, &o.UserID, &o.Name, &o.Address, &o.Styles, &o.Type, &o.Condition, &o.Lng, &o.Lat, &o.Year, &o.ImgCount)
		return &o, err
	})
}

func styleSearch(sb *sqlbuilder.SelectBuilder, s *Search) (string, error) {
	//Style search
	if s.Style == "" {
		return "any", nil
	}

	sty, err := FindStyle(s.Style)
	if err != nil {
		return "", err
	}
	fmt.Println(sty)

	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()
	whereClause.SetFlavor(sqlbuilder.PostgreSQL)

	//returns list of ids in style lineage
	// subquery := fmt.Sprintf("SELECT style_id FROM paths WHERE path ~ '%s'", fmt.Sprintf("*.%d.*", sty.ID))
	subquery := fmt.Sprintf("SELECT name FROM styles JOIN paths ON styles.id = paths.style_id WHERE path ~ '%s'", fmt.Sprintf("*.%d.*", sty.ID))

	// Set the flavor of the WhereClause to PostgreSQL.
	whereClause.SetFlavor(sqlbuilder.PostgreSQL)
	whereClause.AddWhereExpr(cond.Args, cond.Or(cond.In("styles[0]", sqlbuilder.Raw(subquery)), cond.In("styles[1]", sqlbuilder.Raw(subquery)), cond.In("styles[2]", sqlbuilder.Raw(subquery))))
	sb.AddWhereClause(whereClause)

	// a := "we want observations where there is at least 1 overlap between style lineage of search and styles of observation"
	// q := "SELECT * FROM observations
	// 		WHERE EXISTS (SELECT style FROM
	// 		)"

	return "", nil
}

func spatialSearch(sb *sqlbuilder.SelectBuilder, s *Search) string {

	if s.Lat1 == 0 && s.Lng1 == 0 && s.Lat2 == 0 && s.Lng2 == 0 && s.Radius == 0 {
		//No search region defined
		return "any"
	}
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()
	whereClause.SetFlavor(sqlbuilder.PostgreSQL)

	//Spatial search
	if s.Lat1 != 0 && s.Lng1 != 0 && s.Lat2 != 0 && s.Lng2 != 0 && s.Radius == 0 {
		//Square search
		whereClause.AddWhereExpr(
			cond.Args,
			cond.Between(`lng`, s.Lng1, s.Lng2),
			cond.Between(`lat`, s.Lat1, s.Lat2),
		)
		sb.AddWhereClause(whereClause)
	} else if s.Lat1 != 0 && s.Lng1 != 0 && s.Lat2 == 0 && s.Lng2 == 0 && s.Radius != 0 {
		radialSearch(sb, s.Lat1, s.Lng1, s.Radius)
	}

	return ""
}

func temporalSearch(sb *sqlbuilder.SelectBuilder, s *Search) string {

	if s.BefYear == 0 && s.AftYear == 0 {
		return "any"
	}
	whereClause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()
	whereClause.SetFlavor(sqlbuilder.PostgreSQL)

	//Temporal search
	whereClause.AddWhereExpr(
		cond.Args,
		cond.Between(`year`, s.AftYear, s.BefYear),
	)
	cond.Between(`year`, s.AftYear, s.BefYear)

	sb.AddWhereClause(whereClause)
	return ""
}

type ErrBadSearch struct {
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e ErrBadSearch) Error() string {
	return e.Message + ": " + strings.Join(e.Details, ", ")
}

func radialSearch(sb *sqlbuilder.SelectBuilder, lat, lng, radius float64) {
	floatCoords := [3]float64{lng, lat, radius}
	cds := [3]string{}
	for i, v := range floatCoords {
		cds[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}

	region := fmt.Sprintf("(((o.lng-%s)^2 + (o.lat-%s)^2) <= %s^2)", cds[0], cds[1], cds[2])
	// sqlbuilder.Raw(region)

	sb.And(region)
}


