import { createApp } from 'vue'
import App from './App.vue'
import makeFontAwesomePlugin from './plugins/font-awesome'
import { router } from './routes'
import { store } from './store'
import './assets/tailwind.css'

const app = createApp(App)

app.component('font-awesome', makeFontAwesomePlugin())

app.use(store)
app.use(router)
app.mount('#app')
