package main

import (
	edb "eurostat-weekly-deaths/db"
	"eurostat-weekly-deaths/parser"
	"flag"
	"fmt"
	"log"
)

func main() {
	var path string
	flag.StringVar(
		&path,
		"eurostat_input",
		"data/weekly_deaths.tsv.gz",
		"Path to gzip archive with Eurostat data.",
	)
	flag.Parse()

	it, err := parser.NewEurostatWeeklyDeathsData(path)
	if err != nil {
		log.Fatal(err)
	}

	db := edb.DB()
	fmt.Printf("DB created: %+v", db)

	edb.BatchInsertWeeklyDeaths(it, db)
}
