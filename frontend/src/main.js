import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import vuetify from './plugins/vuetify'

Vue.config.productionTip = false
axios.defaults.baseURL = 'http://localhost:8080'
Vue.prototype.$http = axios


new Vue({
  vuetify,
  render: h => h(App)
}).$mount('#app')
