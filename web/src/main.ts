import { createApp } from 'vue'
import App from './App.vue'
import makeFontAwesomePlugin from './plugins/font-awesome'
import { router } from './routes'
import { store } from './store'

import './assets/styles/root.css'

const app = createApp(App)

app.component('FontAwesome', makeFontAwesomePlugin())

app.use(store)
app.use(router)
app.mount('#app')
