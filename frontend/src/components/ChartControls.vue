<template>
  <div>
    <div style="padding: 7px"><v-autocomplete
        :items="countries"
        label="Country"
        item-text="name"
        item-value="code"
        v-model="country"
    >

    </v-autocomplete></div>
    <div style="padding: 7px"><v-select
        :items="genders"
        label="Gender"
        item-text="name"
        item-value="code"
        v-model="gender"
    ></v-select></div>

    <div style="padding: 7px"><v-select
        :items="ages"
        label="Age"
        item-text="name"
        item-value="code"
        v-model="age"
    ></v-select></div>

    <div style="padding: 7px"><v-autocomplete
        :items="years"
        label="Year From"
        item-text="name"
        v-model="yearFrom"
    >
    </v-autocomplete></div>

    <div style="padding: 7px"><v-autocomplete
        :items="years"
        label="Year To"
        item-text="name"
        v-model="yearTo"
    >
    </v-autocomplete></div>
  </div>

</template>

<script>
export default {
  name: "ChartControls",
  props: ['countries', 'ages', 'genders', 'years'],
  emits: ['inputChanged'],
  data: function() {
    return {
      country: this.countries[0].code,
      age: this.ages[0].code,
      gender: this.genders[0].code,
      yearFrom: this.years[0].value,
      yearTo: this.years[this.years.length-1].value,
    }
  },
  computed: {
    queryUrl: function() {
      return `weekly_deaths?country=${this.country}&age=${this.age}&gender=${this.gender}&year_from=${this.yearFrom}&yearTo=${this.yearTo}`
    }
  },
  watch: {
    queryUrl: function(newVal, oldVal) {
      if (newVal !== oldVal && newVal !== undefined) {
        this.$emit('inputChanged', newVal)
      }
    }
  }
}
</script>
