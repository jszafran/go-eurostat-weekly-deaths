package database

import "gorm.io/gorm"

type EurostatRepository struct {
	Db *gorm.DB
}

func (r *EurostatRepository) FetchCountryData(p WeeklyDeathsQueryParams) []WeeklyDeaths {
	filter := "country IN ? and age = ? and gender = ? and year >= ? and year <= ?"

	var res []WeeklyDeaths
	r.Db.Order("year, week").Where(filter, p.Country, p.Age, p.Gender, p.YearFrom, p.YearTo).Find(&res)
	return res
}

func (r *EurostatRepository) FetchFirstN(n int) []WeeklyDeaths {
	var res []WeeklyDeaths

	r.Db.Limit(n).Find(&res)

	return res
}
