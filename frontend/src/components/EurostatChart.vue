<template>
    <div>
      <h3>Line</h3>
      <rough-line
          roughness="8"
          data="https://raw.githubusercontent.com/jwilber/random_data/master/profits.csv"
          y1="revenue"
          y2="cost"
          y3="profit"
      ></rough-line>
    </div>
</template>

<script>
import { RoughLine } from 'vue-roughviz'
export default {
  components: {RoughLine,},
  name: 'EurostatChart',
  props: ['chartData'],
  data: function() {
    return {
    }
  },
}
</script>

<style>

</style>