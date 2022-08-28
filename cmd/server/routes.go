package main

import (
	"eurostat-weekly-deaths/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	country := c.Query("country")
	age := c.Query("age")
	gender := c.Query("gender")

	yearFrom, err := strconv.Atoi(c.Query("yearFrom"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	yearTo, err := strconv.Atoi(c.Query("yearTo"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	res := e.repo.FetchCountryData(country, age, gender, yearFrom, yearTo)
	c.JSON(http.StatusOK, res)
}

func GetRouter(repo *database.EurostatRepository) *gin.Engine {
	r := gin.Default()
	env := Env{repo}
	r.GET("/countries", env.ListCountries)
	r.GET("/ages", env.ListAges)
	r.GET("/genders", env.ListGenders)
	r.GET("/weekly_deaths", env.ListWeeklyDeaths)
	return r
}
