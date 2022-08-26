package parser

import (
	"io"
	"reflect"
	"testing"
)

func TestEurostatWeeklyDeathsDataParser(t *testing.T) {
	// 2022W31 	2022W30 	2022W29
	testRecords := []Record{
		{
			Age:          "Y_LT5",
			Sex:          "F",
			Country:      "PL",
			WeeklyDeaths: 1,
			Year:         2022,
			Week:         31,
		},
		{
			Age:          "Y_LT5",
			Sex:          "F",
			Country:      "PL",
			WeeklyDeaths: 2,
			Year:         2022,
			Week:         30,
		},
		{
			Age:          "Y_LT5",
			Sex:          "F",
			Country:      "PL",
			WeeklyDeaths: -1,
			Year:         2022,
			Week:         29,
		},
		{
			Age:          "Y20-24",
			Sex:          "M",
			Country:      "DE",
			WeeklyDeaths: 100,
			Year:         2022,
			Week:         31,
		},
		{
			Age:          "Y20-24",
			Sex:          "M",
			Country:      "DE",
			WeeklyDeaths: 0,
			Year:         2022,
			Week:         30,
		},
		{
			Age:          "Y20-24",
			Sex:          "M",
			Country:      "DE",
			WeeklyDeaths: -1,
			Year:         2022,
			Week:         29,
		},
		{
			Age:          "TOTAL",
			Sex:          "T",
			Country:      "FR",
			WeeklyDeaths: -1,
			Year:         2022,
			Week:         31,
		},
		{
			Age:          "TOTAL",
			Sex:          "T",
			Country:      "FR",
			WeeklyDeaths: -1,
			Year:         2022,
			Week:         30,
		},
		{
			Age:          "TOTAL",
			Sex:          "T",
			Country:      "FR",
			WeeklyDeaths: -1,
			Year:         2022,
			Week:         29,
		},
	}
	e, err := NewEurostatWeeklyDeathsData("../test_data/eurostat_mock_data.tsv.gz")
	if err != nil {
		t.Fatal(err)
	}

	for i, want := range testRecords {
		got, err := e.Next()
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("(%d): Expected %+v, got %+v", i+1, want, got)
		}
	}

	// test end of iteration
	_, err = e.Next()
	if err != io.EOF {
		t.Fatal("Expected io.EOF error")
	}

}
