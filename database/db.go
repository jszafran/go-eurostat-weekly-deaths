package database

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

func DB() (*gorm.DB, error) {
	var db *gorm.DB
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
		return db, err
	}

	return db, nil
}

func LoadWeeklyDeathsData(path string) error {
	db, err := DB()
	if err != nil {
		return err
	}

	log.Println("Recreating database structure... ")
	db.Exec("DROP TABLE IF EXISTS weekly_deaths;")
	db.AutoMigrate(&WeeklyDeaths{})
	log.Print(" Done!")

	log.Println("Parsing data.")
	it, err := parser.NewEurostatWeeklyDeathsData(path)
	if err != nil {
		return err
	}
	log.Println("Parsing done.")

	log.Println("Starting loading data to database.")
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

		// there are some buggy records in Eurostat data (week 99)
		if r.Week <= 53 {
			wd = append(wd, WeeklyDeaths{
				Age:     r.Age,
				Gender:  r.Gender,
				Country: r.Country,
				Value:   r.WeeklyDeaths,
				Week:    r.Week,
				Year:    r.Year,
			})
		}
	}
	db.CreateInBatches(wd, 5000)
	log.Printf("Inserted %d records in %s.\n", len(wd), time.Since(t1))
	return nil
}
