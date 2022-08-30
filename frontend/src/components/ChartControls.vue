<template>
  <div>
    <div style="padding: 7px"><v-autocomplete
        :items="countries"
        label="Country"
        item-text="name"
        item-value="code"
        v-model="chartChoices.country"
    >

    </v-autocomplete></div>
    <div style="padding: 7px"><v-select
        :items="genders"
        label="Gender"
        item-text="name"
        item-value="code"
        v-model="chartChoices.gender"
    ></v-select></div>

    <div style="padding: 7px"><v-select
        :items="ages"
        label="Age"
        item-text="name"
        item-value="code"
        v-model="chartChoices.age"
    ></v-select></div>

    <div style="padding: 7px"><v-autocomplete
        :items="years"
        label="Year From"
        item-text="name"
        v-model="chartChoices.yearFrom"
    >
    </v-autocomplete></div>

    <div style="padding: 7px"><v-autocomplete
        :items="years"
        label="Year To"
        item-text="name"
        v-model="chartChoices.yearTo"
    >
    </v-autocomplete></div>
  </div>

</template>

<script>
export default {
  name: "ChartControls",
  props: ['countries', 'ages', 'genders', 'years', 'chartControlChoice'],
  emits: ['inputChanged'],
  data: function() {
    return {
      chartChoices: this.chartControlChoice,
    }
  },
  watch: {
    chartChoices:
        {
          handler(newVal) {
            this.$emit('inputChanged', newVal)
          },
          deep: true
        }
    }
}
</script>
