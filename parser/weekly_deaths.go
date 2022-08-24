package parser

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

type Record struct {
	MetaData
	WeeklyDeaths
}

type MetaData struct {
	Age     string
	Sex     string
	Country string
}

type WeeklyDeaths struct {
	Year  int
	Week  int
	Value int
}

func parseMetaData(l string) MetaData {
	spl := strings.Split(l, ",")
	return MetaData{
		Age:     spl[0],
		Sex:     spl[1],
		Country: spl[3],
	}
}

func parseWeeklyDeaths(recs []string) []WeeklyDeaths {
	var wd []WeeklyDeaths
	return wd
}

func weekColsMap(rcs []string) map[int]string {
	m := make(map[int]string)

	for i, col := range rcs {
		if i == 0 {
			continue
		}
		m[i] = col
	}
	return m
}

func readGzippedTSV(path string) (*csv.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(gz)
	r.Comma = '\t'
	return r, nil
}

func Parse(path string) ([]Record, error) {
	var recs []Record
	r, err := readGzippedTSV(path)
	if err != nil {
		return recs, err
	}

	hdr, err := r.Read()
	fmt.Println(hdr[0])
	if err != nil {
		return recs, err
	}
	//_ := weekColsMap(hdr)

	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
				fmt.Println("End of file reached!")
				break
			}
		}
		meta := parseMetaData(rec[0])
		fmt.Printf("%+v\n", meta)
	}

	return recs, err
}
