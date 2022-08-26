package db

import (
	"eurostat-weekly-deaths/parser"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type WeeklyDeaths struct {
	gorm.Model
	Age     string
	Sex     string
	Country string
	Value   int
	Week    int
	Year    int
}

func DB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("DROP TABLE IF EXISTS weekly_deaths;")
	db.AutoMigrate(&WeeklyDeaths{})
	return db
}

func BatchInsertWeeklyDeaths(it *parser.EurostatWeeklyDeathsData, db *gorm.DB) error {
	for i := 0; i < 10000; i++ {
		wd, err := it.Next()
		if err != nil {
			return err
		}
		db.Create(&WeeklyDeaths{
			Age:     wd.Age,
			Sex:     wd.Sex,
			Country: wd.Country,
			Value:   wd.WeeklyDeaths,
			Week:    wd.Week,
			Year:    wd.Year,
		})
	}
	return nil
}
