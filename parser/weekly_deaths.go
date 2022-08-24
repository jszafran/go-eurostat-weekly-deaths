package parser

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
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
	if err != nil {
		return recs, err
	}
	wcm := weekColsMap(hdr)
	fmt.Println(wcm)
	return recs, err
}
