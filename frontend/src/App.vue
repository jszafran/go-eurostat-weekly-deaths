<template>
  <v-app id="inspire">
    <v-app-bar
        app
        color="white"
        flat
    >
      <v-container class="py-0 fill-height">
        <v-avatar
            class="mr-10"
            color="grey darken-1"
            size="32"
        ></v-avatar>

        <v-btn
            v-for="link in links"
            :key="link"
            text
        >
          {{ link }}
        </v-btn>
        <v-spacer></v-spacer>

      </v-container>
    </v-app-bar>

    <v-main class="grey lighten-3">
      <v-container>
        <v-row>
          <v-col cols="3">
            <v-sheet rounded="lg">
              <v-list color="transparent">
                <div v-if="dataLoaded">
                  <ChartControls
                      :countries="countries"
                      :ages="ages"
                      :genders="genders"
                      :years="years"
                      @inputChanged="onInputChanged"
                  ></ChartControls>

                </div>


              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet
                min-height="70vh"
                rounded="lg"
            >
              <EurostatChart
                  :gender="gender"
                  :country="country"
                  :age="age"
                  :years="years"
              ></EurostatChart>

            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import EurostatChart from "@/components/EurostatChart";
import ChartControls from "@/components/ChartControls";
export default {
  components: {EurostatChart, ChartControls},
  data: () => ({
    links: [
      'Dashboard',
      'About',
    ],
    country: null,
    gender: null,
    age:  null,
    countries: null,
    genders: null,
    ages: null,
    years: null,
    yearFrom: null,
    yearTo: null,
    chartDataUrl: null,
    dataLoaded: false
  }),
  methods: {
    fetchDropdownValues: async function() {
      const [genders, ages, countries, availableYears] = await Promise.all([
          this.$http.get("/genders"),
          this.$http.get("/ages"),
          this.$http.get("/countries"),
          this.$http.get("/available_years")
      ])
      return {
        genders: genders.data,
        ages: ages.data,
        countries: countries.data,
        yearFrom: availableYears.data.year_from,
        yearTo: availableYears.data.year_to
      }
    },
    fetchChartDataForUrl: async function(url) {
      const resp = await this.$http.get(url)
      return resp.data
    },
    onInputChanged: async function(v) {
      console.log(`Should fetch data for ${v}`)
    }
  },
  mounted: async function() {
    const data = await this.fetchDropdownValues()
    this.genders = data.genders
    this.countries = data.countries
    this.ages = data.ages
    this.country = data.countries[0].code
    this.age = data.ages[0].code
    this.gender = data.genders[0].code

    const years = Array()
    for (let i = data.yearFrom; i <= data.yearTo; i++) {
      years.push({name: i, value: i})
    }
    this.years = years
    this.yearFrom = years[0].name
    this.yearTo = years[2].name
    this.chartDataUrl = `/weekly_deaths?country=${this.country}&age=${this.age}&gender=${this.gender}&year_from=${this.yearFrom}&year_to=${this.yearTo}`
    this.dataLoaded = true
  }
}
</script>