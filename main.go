package main

import (
	"eurostat-weekly-deaths/parser"
	"log"
)

func main() {
	_, err := parser.Parse("data/weekly_deaths.tsv")
	if err != nil {
		log.Fatal(err)
	}
}
