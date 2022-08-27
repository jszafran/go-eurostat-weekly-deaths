package db

import "gorm.io/gorm"

type WeeklyDeaths struct {
	gorm.Model
	Age     string `gorm:"index",json:"age"`
	Gender  string `gorm:"index",json:"gender"`
	Country string `gorm:"index",json:"country"`
	Value   int    `json:"value"`
	Week    int    `gorm:"index",json:"week"`
	Year    int    `gorm:"index",json:"year""`
}

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Gender struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Age struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
