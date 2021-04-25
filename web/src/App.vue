<script lang="ts">
import { defineComponent, resolveComponent, h, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from './store'
import axios from './utils/axios'
import { ActionTypes as AuthActionTypes } from './modules/auth/store/actions'
import { ActionTypes as ThemeActionTypes } from './modules/theme/store/actions'
import { Theme } from './modules/theme/types'

export default defineComponent({
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
      document.querySelector('body')?.classList.add('bg-white', 'dark:bg-black')
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
        () => store.getters['currentTheme'],
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
</script>
