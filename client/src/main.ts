import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import 'vue2-dropzone/dist/vue2Dropzone.min.css'
import VueClipboard from 'vue-clipboard2'

Vue.use(VueClipboard)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
