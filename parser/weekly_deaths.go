package parser

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Age          string
	Sex          string
	Country      string
	WeeklyDeaths int
	Year         int
	Week         int
}

type Demographics struct {
	Age     string
	Sex     string
	Country string
}

type WeeklyDeaths struct {
	Value int
}

type WeekInfo struct {
	Year int
	Week int
}

func parseDemographics(l string) Demographics {
	spl := strings.Split(l, ",")
	return Demographics{
		Age:     spl[0],
		Sex:     spl[1],
		Country: spl[3],
	}
}

func parseWeeklyDeathsValue(v string) (int, error) {
	var vi int

	v = strings.Replace(v, " ", "", -1)

	if strings.Contains(v, ":") {
		return -1, nil
	}

	if strings.Contains(v, "p") {
		v = strings.Replace(v, "p", "", -1)
	}

	vi, err := strconv.Atoi(v)
	if err != nil {
		return vi, err
	}
	return vi, nil
}

func parseWeeklyDeathsLine(rec []string) ([]int, error) {
	wd := make([]int, len(rec))
	for i, r := range rec {
		v, err := parseWeeklyDeathsValue(r)
		if err != nil {
			return wd, err
		}
		wd[i] = v
	}
	return wd, nil
}

func parseWeekInfo(s string) (WeekInfo, error) {
	var wi WeekInfo

	s = strings.Replace(s, " ", "", -1)

	spl := strings.Split(s, "W")

	if len(spl) != 2 {
		return wi, fmt.Errorf("bad header value: %s", s)
	}

	yr, err := strconv.Atoi(spl[0])
	if err != nil {
		return wi, err
	}

	wk, err := strconv.Atoi(spl[1])
	if err != nil {
		return wi, err
	}

	return WeekInfo{Week: wk, Year: yr}, nil
}

func weekInfoColsMap(r []string) (map[int]WeekInfo, error) {
	m := make(map[int]WeekInfo)

	for i, col := range r {
		if i == 0 {
			continue
		}

		wi, err := parseWeekInfo(col)
		if err != nil {
			return m, err
		}

		m[i] = WeekInfo{
			Year: wi.Year,
			Week: wi.Week,
		}
	}
	return m, nil
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
	wim, err := weekInfoColsMap(hdr)

	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		}
		dmg := parseDemographics(rec[0])
		vals, err := parseWeeklyDeathsLine(rec[1:])
		if err != nil {
			return recs, err
		}

		for i, v := range vals {
			wi := wim[i+1]
			recs = append(recs, Record{
				Age:          dmg.Age,
				Sex:          dmg.Sex,
				Country:      dmg.Country,
				WeeklyDeaths: v,
				Year:         wi.Year,
				Week:         wi.Week,
			})
		}
	}

	return recs, err
}

type SourceIterator struct {
	reader       *csv.Reader
	row          []string
	rowPos       int
	demographics *Demographics
	wim          map[int]WeekInfo
	lastRow      bool
}

func (si *SourceIterator) Next() (Record, error) {
	var rec Record
	// set demographics and row
	if si.rowPos == 0 {
		row, err := si.reader.Read()
		if err != nil {
			if err == io.EOF {
				return rec, io.EOF
			} else {
				return rec, err
			}
		}
		si.row = row
		si.rowPos++
		dmg := parseDemographics(si.row[0])
		si.demographics = &dmg
	}

	// parse values from row
	v, err := parseWeeklyDeathsValue(si.row[si.rowPos])
	if err != nil {
		return rec, err
	}

	wi, exists := si.wim[si.rowPos]
	if !exists {
		return rec, fmt.Errorf("couldn't find a week info for position %d", si.rowPos)
	}

	// TODO: prepare for next row
	if si.rowPos+1 == len(si.row) && si.lastRow {
		return rec, io.EOF
	} else if si.rowPos+1 == len(si.row) {
		si.rowPos = 0
	} else {
		si.rowPos++
	}

	return Record{
		Age:          si.demographics.Age,
		Sex:          si.demographics.Sex,
		Country:      si.demographics.Country,
		WeeklyDeaths: v,
		Year:         wi.Year,
		Week:         wi.Week,
	}, nil
}

func NewEurostatIterator(path string) (*SourceIterator, error) {
	var si SourceIterator
	r, err := readGzippedTSV(path)
	if err != nil {
		return &si, err
	}

	hdr, err := r.Read()
	if err != nil {
		return &si, err
	}

	wim, err := weekInfoColsMap(hdr)
	if err != nil {
		return &si, err
	}

	return &SourceIterator{
		reader:       r,
		row:          nil,
		rowPos:       0,
		demographics: nil,
		wim:          wim,
	}, nil

}
