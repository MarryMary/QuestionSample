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
import axios from "axios";

library.add(fas, far, fab)

axios.get(
    'http://localhost:8080/IsAuthMe',
).then(
    function(response){
        if(response.status === 200){
            var data = response.data
            if(data.AUTH_STATUS_NUMBER !== "1"){
                location.href='http://localhost/AuthSample/Auth/login.php'
            }
        }else{
            location.href='http://localhost/AuthSample/Auth/login.php'
        }
    }.bind(this)
).catch(
    error => console.log(error)
)

var app = createApp(App)
app.component('font-awesome-icon', FontAwesomeIcon)
app.use(VueMathjax)
app.use(store)
app.use(router)
app.mount('#app')

