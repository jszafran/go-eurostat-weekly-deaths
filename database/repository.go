package database

import "gorm.io/gorm"

type EurostatRepository struct {
	Db *gorm.DB
}

func (r *EurostatRepository) FetchCountryData(
	country string,
	age string,
	gender string,
	yearFrom int,
	yearTo int,
) []WeeklyDeaths {
	filter := "country = ? and age = ? and gender = ? and year >= ? and year <= ?"

	var res []WeeklyDeaths
	r.Db.Order("year, week").Where(filter, country, age, gender, yearFrom, yearTo).Find(&res)
	return res
}

func (r *EurostatRepository) FetchFirstN(n int) []WeeklyDeaths {
	var res []WeeklyDeaths

	r.Db.Limit(n).Find(&res)

	return res
}
