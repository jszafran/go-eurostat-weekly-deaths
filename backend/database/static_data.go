package database

import "eurostat-weekly-deaths/models"

func Countries() []models.Country {
	return []models.Country{
		{"Andorra", "AD"},
		{"Albania", "AL"},
		{"Armenia", "AM"},
		{"Austria", "AT"},
		{"Belgium", "BE"},
		{"Bulgaria", "BG"},
		{"Switzerland", "CH"},
		{"Cyprus", "CY"},
		{"Czechia", "CZ"},
		{"Germany", "DE"},
		{"Denmark", "DK"},
		{"Estonia", "EE"},
		{"Greece", "EL"},
		{"Spain", "ES"},
		{"Finland", "FI"},
		{"France", "FR"},
		{"Georgia", "GE"},
		{"Croatia", "HR"},
		{"Hungary", "HU"},
		{"Ireland", "IE"},
		{"Iceland", "IS"},
		{"Italy", "IT"},
		{"Liechtenstein", "LI"},
		{"Lithuania", "LT"},
		{"Luxembourg", "LU"},
		{"Latvia", "LV"},
		{"Montenegro", "ME"},
		{"Malta", "MT"},
		{"Netherlands", "NL"},
		{"Norway", "NO"},
		{"Poland", "PL"},
		{"Portugal", "PT"},
		{"Romania", "RO"},
		{"Serbia", "RS"},
		{"Sweden", "SE"},
		{"Slovenia", "SI"},
		{"Slovakia", "SK"},
		{"United Kingdom", "UK"},
	}
}

func Genders() []models.Gender {
	return []models.Gender{
		{"Total", "T"},
		{"Female", "F"},
		{"Male", "M"},
	}
}

func Ages() []models.Age {
	return []models.Age{
		{"Total", "TOTAL"},
		{"<5 years old", "Y_LT5"},
		{"5-9 years old", "Y5-9"},
		{"10-14 years old", "Y10-14"},
		{"15-19 years old", "Y15-19"},
		{"20-24 years old", "Y20-24"},
		{"25-29 years old", "Y25-29"},
		{"30-34 years old", "Y30-34"},
		{"35-39 years old", "Y35-39"},
		{"40-44 years old", "Y40-44"},
		{"45-49 years old", "Y45-49"},
		{"50-54 years old", "Y50-54"},
		{"55-59 years old", "Y55-59"},
		{"60-64 years old", "Y60-64"},
		{"65-69 years old", "Y65-69"},
		{"70-74 years old", "Y70-74"},
		{"75-79 years old", "Y75-79"},
		{"80-84 years old", "Y80-84"},
		{"85-89 years old", "Y85-89"},
		{">=90 years old", "Y_GE90"},
	}
}
