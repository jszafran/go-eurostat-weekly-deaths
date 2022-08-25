package main

import (
	"eurostat-weekly-deaths/parser"
	"fmt"
	"log"
	"time"
)

func main() {
	st := time.Now()
	r, err := parser.Parse("data/weekly_deaths.tsv")
	if err != nil {
		log.Fatal(err)
	}
	tt := time.Since(st)
	fmt.Printf("Parsed %d records in %v.", len(r), tt)
}
