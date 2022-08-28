package main

import (
	"eurostat-weekly-deaths/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCountries(c *gin.Context) {
	c.JSON(http.StatusOK, db.Countries())
}

func ListGenders(c *gin.Context) {
	c.JSON(http.StatusOK, db.Genders())
}

func ListAges(c *gin.Context) {
	c.JSON(http.StatusOK, db.Ages())
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/countries", ListCountries)
	r.GET("/ages", ListAges)
	r.GET("/genders", ListGenders)
	return r
}
