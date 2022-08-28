package main

import (
	"eurostat-weekly-deaths/database"
	"log"
)

func main() {
	db, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}

	repo := database.EurostatRepository{Db: db}
	r := GetRouter(&repo)
	log.Fatal(r.Run("localhost:8080"))
}
