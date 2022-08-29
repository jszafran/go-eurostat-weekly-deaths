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
                <div style="padding: 7px"><v-autocomplete
                    :items="countries"
                    label="Country"
                    item-text="name"
                    v-model="country.value"
                >

                </v-autocomplete></div>
                <div style="padding: 7px"><v-select
                    :items="genders"
                    label="Gender"
                    item-text="name"
                    v-model="gender.value"
                ></v-select></div>


                <v-divider class="my-2"></v-divider>
                <v-list-item
                    link
                    color="grey lighten-4"
                >
                  <v-list-item-content>
                    <v-list-item-title>
                      Refresh
                    </v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet
                min-height="70vh"
                rounded="lg"
            >
              <EurostatChart :gender="gender.value" :country="country.value"></EurostatChart>
              <!--  -->
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import EurostatChart from "@/components/EurostatChart";

export default {
  components: {EurostatChart},
  data: () => ({
    links: [
      'Dashboard',
      'About',
      'Fo'
    ],
    country: {name: 'Poland', 'value': 'PL'},
    gender: {name: 'All', 'value': 'TOTAL'},
    countries: [
      {'name': 'Poland', 'value': 'PL'},
      {'name': 'Germany', 'value': 'DE'},
      {'name': 'France', 'value': 'FR'}
    ],
    genders: [
      {'name': 'All', 'value': 'TOTAL'},
      {'name': 'Female', 'value': 'Female'},
      {'name': 'Male', 'value': 'Male'}
    ]
  }),
  methods: {
    fetch
  }
}
</script>