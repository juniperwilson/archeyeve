package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juniperwilson/archeyeve/internal/database"
)

func createObservations(c *gin.Context) {
	var obs database.Observation
	if err := c.BindJSON(&obs); err != nil {
		c.Error(apiError{[]string{"bad input"}})
		return
	}

	errs := []string{}

	if obs.Name == "" {
		errs = append(errs, "must have a name")
	}

	if obs.Lng == 0.0 && obs.Lat == 0.0 {
		errs = append(errs, "invalid location")
	}

	if obs.Year > time.Now().Year() {
		errs = append(errs, "year out of bounds")
	}

	if obs.ImgCount == 0 {
		errs = append(errs, "must have an image")
	}

	if obs.Styles == nil {
		errs = append(errs, "must have a style")
	}

	_, err := database.AddObservation(&obs)
	if err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) != 0 {
		c.Error(apiError{errs})
		return
	}
}

func search(c *gin.Context) {
	var search database.Search
	errs := []string{}

	search.Style = c.Query("style")

	if lng := c.Query("lng1"); lng != "" {
		lng1, err := strconv.ParseFloat(lng, 64)
		if err != nil {
			errs = append(errs, err.Error())
		}
		search.Lng1 = lng1
	}
	if lat := c.Query("lat1"); lat != "" {
		lat1, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			errs = append(errs, err.Error())
		}
		search.Lat1 = lat1
	}
	if lng := c.Query("lng2"); lng != "" {
		lng2, err := strconv.ParseFloat(lng, 64)
		if err != nil {
			errs = append(errs, err.Error())
		}
		search.Lng2 = lng2
	}
	if lat := c.Query("lat2"); lat != "" {
		lat2, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			errs = append(errs, err.Error())
		}
		search.Lat2 = lat2
	}
	if r := c.Query("radius"); r != "" {
		radius, err := strconv.ParseFloat(r, 64)
		if err != nil {
			errs = append(errs, err.Error())
		}
		search.Radius = radius
	}
	if by := c.Query("befyear"); by != "" {
		befyear, err := strconv.Atoi(by)
		if err != nil {
			errs = append(errs, err.Error())
		}
		search.BefYear = befyear
	}
	if ay := c.Query("aftyear"); ay != "" {
		aftyear, err := strconv.Atoi(ay)
		if err != nil {
			errs = append(errs, err.Error())
		}
		search.AftYear = aftyear
	}

	if len(errs) != 0 {
		c.Error(apiError{errs})
		return
	}

	errs = []string{}

	fmt.Println(search)

	obs, err := database.Find(&search)
	if err != nil {
		fmt.Println(err.Error())
		errs = append(errs, "internal error")
	}

	fmt.Println(obs)

	if len(errs) != 0 {
		c.Error(apiError{errs})
		return
	}

	c.JSON(http.StatusOK, obs)
}

func ApiRouting() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/search", search)
	router.POST("/observation", createObservations)

	return router
}
