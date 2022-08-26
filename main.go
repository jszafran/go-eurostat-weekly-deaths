package main

import (
	edb "eurostat-weekly-deaths/db"
	"eurostat-weekly-deaths/parser"
	"flag"
	"fmt"
	"log"
	"time"
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
	
	t1 := time.Now()
	edb.BatchInsertWeeklyDeaths(it, db)
	fmt.Printf("Inserted records. Time taken: %s", time.Since(t1))

}
