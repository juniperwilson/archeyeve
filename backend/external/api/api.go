package api

import (
	"net/http"
	"strconv"
	"time"

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
	lng1, err := strconv.ParseFloat(c.Query("lng1"), 64)
	if err != nil {
		errs = append(errs, err.Error())
	}
	search.Lng1 = lng1
	lat1, err := strconv.ParseFloat(c.Query("lat1"), 64)
	if err != nil {
		errs = append(errs, err.Error())
	}
	search.Lat1 = lat1
	lng2, err := strconv.ParseFloat(c.Query("lng2"), 64)
	if err != nil {
		errs = append(errs, err.Error())
	}
	search.Lng2 = lng2
	lat2, err := strconv.ParseFloat(c.Query("lat2"), 64)
	if err != nil {
		errs = append(errs, err.Error())
	}
	search.Lat2 = lat2
	radius, err := strconv.ParseFloat(c.Query("radius"), 64)
	if err != nil {
		errs = append(errs, err.Error())
	}
	search.Radius = radius
	befyear, err := strconv.Atoi(c.Query("befyear"))
	if err != nil {
		errs = append(errs, err.Error())
	}
	search.BefYear = befyear
	aftyear, err := strconv.Atoi(c.Query("aftyear"))
	if err != nil {
		errs = append(errs, err.Error())
	}
	search.AftYear = aftyear
	if len(errs) != 0 {
		c.Error(apiError{errs})
		return
	}

	errs = []string{}

	obs, err := database.Find(&search)
	if err != nil {
		errs = append(errs, "internal error")
	}

	if len(errs) != 0 {
		c.Error(apiError{errs})
		return
	}

	c.JSON(http.StatusOK, obs)
}

func ApiRouting() *gin.Engine {
	router := gin.Default()

	router.GET("/search", search)
	router.POST("/observation", createObservations)

	return router
}
