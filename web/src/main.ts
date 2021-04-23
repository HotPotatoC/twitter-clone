import {
  createApp,
  defineComponent,
  resolveComponent,
  h,
  watch,
  onMounted,
} from 'vue'
import { useRouter } from 'vue-router'

import { router } from './routes'
import { store, useStore } from './store'
import axios from './utils/axios'
import { ActionTypes as AuthActionTypes } from './modules/auth/store/actions'
import { ActionTypes as ThemeActionTypes } from './modules/theme/store/actions'

import './assets/styles/root.css'
import { Theme } from './modules/theme/types'

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

            await store.dispatch(AuthActionTypes.REFRESH_AUTH_TOKEN)
            await store.dispatch(AuthActionTypes.GET_USER_DATA)
            if (!store.getters['isLoggedIn']) {
              router.push('/login')
              return Promise.reject(error)
            }
          }
        }
      )

      onMounted(() => {
        document
          .querySelector('body')
          ?.classList.add('bg-white', 'dark:bg-black')
        if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
          window.localStorage.setItem('theme', 'dark')
          store.dispatch(ThemeActionTypes.TOGGLE_THEME, 'dark')
          document.querySelector('html')?.classList.add('dark')
        } else {
          window.localStorage.setItem('theme', 'dark')
          store.dispatch(ThemeActionTypes.TOGGLE_THEME, 'light')
          document.querySelector('html')?.classList.remove('dark')
        }

        watch(
          () => window.localStorage.getItem('theme'),
          (theme) => {
            store.dispatch(ThemeActionTypes.TOGGLE_THEME, theme as Theme)

            if (store.getters['currentTheme'] === 'dark') {
              document.querySelector('html')?.classList.add('dark')
            } else {
              document.querySelector('html')?.classList.remove('dark')
            }
          }
        )
      })

      return () => h(RouterView)
    },
  })
)

app.use(store)
app.use(router)
app.mount('#app')
