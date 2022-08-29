package main

import (
	"eurostat-weekly-deaths/database"
	"eurostat-weekly-deaths/models"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Env struct {
	repo *database.EurostatRepository
}

func (e *Env) ListCountries(c *gin.Context) {
	c.JSON(http.StatusOK, database.Countries())
}

func (e *Env) ListGenders(c *gin.Context) {
	c.JSON(http.StatusOK, database.Genders())
}

func (e *Env) ListAges(c *gin.Context) {
	c.JSON(http.StatusOK, database.Ages())
}

func (e *Env) ListWeeklyDeaths(c *gin.Context) {
	if len(c.Request.URL.Query()) == 0 {
		c.JSON(http.StatusOK, e.repo.FetchFirstN(20))
		return
	}

	var params models.WeeklyDeathsQueryParams
	err := c.Bind(&params)
	fmt.Printf("%+v\n", params)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res := e.repo.FetchCountryData(params)
	c.JSON(http.StatusOK, res)
}

func GetRouter(repo *database.EurostatRepository) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	env := Env{repo}
	r.GET("/countries", env.ListCountries)
	r.GET("/ages", env.ListAges)
	r.GET("/genders", env.ListGenders)
	r.GET("/weekly_deaths", env.ListWeeklyDeaths)
	return r
}
