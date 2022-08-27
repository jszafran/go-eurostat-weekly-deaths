package db

import (
	"eurostat-weekly-deaths/parser"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"time"
)

type WeeklyDeaths struct {
	gorm.Model
	Age     string
	Gender  string
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
	t1 := time.Now()
	wd := make([]WeeklyDeaths, 0)
	for {
		r, err := it.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		wd = append(wd, WeeklyDeaths{
			Age:     r.Age,
			Gender:  r.Gender,
			Country: r.Country,
			Value:   r.WeeklyDeaths,
			Week:    r.Week,
			Year:    r.Year,
		})
	}
	db.CreateInBatches(wd, 5000)
	log.Printf("Inserted %d records in %s.\n", len(wd), time.Since(t1))
	return nil
}
