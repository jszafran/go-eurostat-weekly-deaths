package database

import (
	"eurostat-weekly-deaths/models"
	"gorm.io/gorm"
)

type EurostatRepository struct {
	Db *gorm.DB
}

func (r *EurostatRepository) FetchCountryData(p models.WeeklyDeathsQueryParams) []models.WeeklyDeaths {
	var res []models.WeeklyDeaths

	r.Db.Order(
		"year, week",
	).Where(
		"country = ? and age = ? and gender = ? and year >= ? and year <= ?",
		p.Country,
		p.Age,
		p.Gender,
		p.YearFrom,
		p.YearTo,
	).Find(&res)

	return res
}

func (r *EurostatRepository) FetchFirstN(n int) []models.WeeklyDeaths {
	var res []models.WeeklyDeaths

	r.Db.Limit(n).Find(&res)

	return res
}

func (r *EurostatRepository) GetAvailableYears() models.AvailableYears {
	var res models.AvailableYears

	r.Db.Raw("SELECT min(year) AS year_from, max(year) AS year_to FROM weekly_deaths;").Scan(&res)

	return res
}
