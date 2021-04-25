import {
  createApp,
} from 'vue'

import { router } from './routes'
import { store } from './store'

import './assets/styles/root.css'
import App from './App.vue'

const app = createApp(App)

app.use(store)
app.use(router)
app.mount('#app')
