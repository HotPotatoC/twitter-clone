import { createApp, defineComponent, resolveComponent, h } from 'vue'
import { useRouter } from 'vue-router'
import { router } from './routes'
import { store, useStore } from './store'
import axios from './services/axios'
import { ActionTypes } from './modules/auth/store/actions'
import makeFontAwesomePlugin from './plugins/font-awesome'

import './assets/styles/root.css'

const app = createApp(
  defineComponent({
    setup() {
      const store = useStore()
      const router = useRouter()
      const RouterView = resolveComponent('router-view')

      const responseInterceptor = axios.interceptors.response.use(
        (res) => res,
        async (error) => {
          if (error.response.status === 401) {
            axios.interceptors.response.eject(responseInterceptor)

            await store.dispatch(ActionTypes.REFRESH_AUTH_TOKEN)
            await store.dispatch(ActionTypes.GET_USER_DATA)
            if (!store.getters['isLoggedIn']) {
              router.push('/login')
              return Promise.reject(error)
            }
          }
        }
      )

      return () => h(RouterView)
    },
  })
)

app.component('FontAwesome', makeFontAwesomePlugin())

app.use(store)
app.use(router)
app.mount('#app')
