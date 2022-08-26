package main

import (
	"eurostat-weekly-deaths/db"
	"eurostat-weekly-deaths/parser"
	"flag"
	"fmt"
	"io"
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

	t1 := time.Now()
	it, err := parser.NewEurostatWeeklyDeathsData(path)
	if err != nil {
		log.Fatal(err)
	}

	rcs := make([]parser.Record, 0)
	for {
		r, err := it.Next()
		rcs = append(rcs, r)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
	}

	fmt.Printf("Parsed %d records in %s.", len(rcs), time.Since(t1))

	db := db.DB()
	fmt.Printf("DB created: %+v", db)
}
