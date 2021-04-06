<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'
import axios from './services/axios'
import { useStore } from './store'
import { ActionTypes } from './store/auth/actions'

export default defineComponent({
  name: 'App',
  setup() {
    const store = useStore()
    const router = useRouter()

    store
      .dispatch(ActionTypes.REFRESH_AUTH_TOKEN)
      .then(() => {
        if (!store.getters['isLoggedIn']) {
          router.push('/login')
        }
      })
      .then(async () => {
        await store.dispatch(ActionTypes.GET_USER_DATA)
        if (!store.getters['isLoggedIn']) {
          router.push('/login')
        }
      })

    const interceptor = axios.interceptors.response.use(
      (res) => {
        return res
      },
      async (error) => {
        if (error.response.status === 401) {
          axios.interceptors.response.eject(interceptor)

          await store.dispatch(ActionTypes.REFRESH_AUTH_TOKEN)
          await store.dispatch(ActionTypes.GET_USER_DATA)
          if (!store.getters['isLoggedIn']) {
            router.push('/login')
            return Promise.reject(error)
          }
        }
      }
    )
  },
})
</script>

<style>
@import url('https://rsms.me/inter/inter.css');
html {
  font-family: 'Inter', sans-serif;
}
@supports (font-variation-settings: normal) {
  html {
    font-family: 'Inter var', sans-serif;
  }
}

@media (prefers-color-scheme: dark) {
  body {
    --tw-bg-opacity: 1;
    background-color: rgba(0, 0, 0, var(--tw-bg-opacity));
  }
}

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.hide-scrollbar::-webkit-scrollbar {
  width: 0px;
  background: transparent; /* Chrome/Safari/Webkit */
}

.hide-scrollbar {
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none; /* IE 10+ */
}
</style>
