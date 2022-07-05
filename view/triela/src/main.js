import'bootstrap/dist/css/bootstrap.css'
import { createApp } from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import VueMathjax from 'vue-mathjax-next'
import App from './App.vue'
import router from './router'
import store from './store'

library.add(fas, far, fab)

var app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon)
app.use(VueMathjax)
app.use(store)
app.use(router)
app.mount('#app')

