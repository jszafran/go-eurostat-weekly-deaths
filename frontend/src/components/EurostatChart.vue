<template>
  <div><p>Selected combination: {{ country }}; {{ gender }}; {{ age }}</p></div>
</template>

<script>
export default {
  name: "EurostatChart",
  props: ['country', 'gender', 'age'],
  data: function() {
    return {
      opts: Array("One", "Two", "Three")
    }
  }
}
</script>

<style>

</style>