import { createApp } from 'vue'
//import App from './App.vue'
import Triela from './Triela.vue'
import router from './router'
import store from './store'
import "bootstrap/dist/css/bootstrap.min.css"

createApp(Triela).use(store).use(router).mount('#app')
