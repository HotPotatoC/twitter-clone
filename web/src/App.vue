<script lang="ts">
import { defineComponent, resolveComponent, h } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from './store'
import axios from './services/axios'
import { ActionTypes } from './modules/auth/store/actions'

export default defineComponent({
  name: 'App',
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
</script>
