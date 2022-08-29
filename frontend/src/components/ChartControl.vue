<template>
  <v-select items="opts" label="Country"></v-select>
</template>

<script>
export default {
  name: "ChartControl",
  data: function() {
    return {
      opts: Array("One", "Two", "Three")
    }
  }
}
</script>

<style>

</style>