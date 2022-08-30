export class ChartControlChoice {
    constructor(country, age, gender, yearFrom, yearTo) {
        this.country = country
        this.age = age
        this.gender = gender
        this.yearFrom = yearFrom
        this.yearTo = yearTo
    }

    buildUrl() {
        return `weekly_deaths?country=${this.country}&age=${this.age}&gender=${this.gender}&year_from=${this.yearFrom}&year_to=${this.yearTo}`
    }
}

