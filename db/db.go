package db

import (
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
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("DROP TABLE IF EXISTS weekly_deaths;")
	db.AutoMigrate(&WeeklyDeaths{})
	return db
}
