package main

import (
	appdb "eurostat-weekly-deaths/db"
	"flag"
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

	err := appdb.LoadWeeklyDeathsData(path)
	if err != nil {
		log.Fatal(err)
	}
}
