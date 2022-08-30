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
                      :chart-control-choice="chartControls"
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
                :chart-data="chartData"
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
import {ChartControlChoice} from "@/chart_choices.js"

export default {
  components: {EurostatChart, ChartControls},
  data: () => ({
    links: [
      'Dashboard',
      'About',
    ],
    countries: null,
    genders: null,
    ages: null,
    years: null,
    yearFrom: null,
    yearTo: null,
    dataLoaded: false,
    chartData: null,
    chartControls: null
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

    const years = Array()
    for (let i = data.yearFrom; i <= data.yearTo; i++) {
      years.push({name: i, value: i})
    }
    this.years = years

    this.chartControls = new ChartControlChoice(
        data.countries[0].code,
        data.ages[0].code,
        data.genders[0].code,
        years[0].value,
        years[years.length - 1].value
    )
    this.dataLoaded = true
  }
}
</script>