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
  data: function() {
    return {
      country: this.countries[0],
      age: this.ages[0],
      gender: this.genders[0],
      yearFrom: this.years[0],
      yearTo: this.years[0]
    }
  },
}
</script>
