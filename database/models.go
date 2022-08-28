package database

type WeeklyDeaths struct {
	Age     string `gorm:"index" json:"age"`
	Gender  string `gorm:"index" json:"gender"`
	Country string `gorm:"index" json:"country"`
	Value   int    `json:"value"`
	Week    int    `gorm:"index" json:"week"`
	Year    int    `gorm:"index" json:"year""`
}

type WeeklyDeathsQueryParams struct {
	Age      string   `form:"age"`
	Gender   string   `form:"gender"`
	Country  []string `form:"country"`
	YearFrom int      `form:"year_from"`
	YearTo   int      `form:"year_to""`
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
