<script lang="ts">
import { computed, defineComponent, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import IconTwitter from '../../icons/IconTwitter.vue'
import LoadingSpinner from '../../shared/LoadingSpinner.vue'
import { useStore } from '../../store'
import { Action } from '../storeActionTypes'

export default defineComponent({
  name: 'Login',
  components: {
    IconTwitter,
    LoadingSpinner,
  },
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()

    const email = ref('')
    const password = ref('')
    const loading = ref(false)

    const input = reactive({ email, password })

    const inputEmpty = computed(
      () => input.email === '' || input.password === ''
    )

    async function authenticate() {
      loading.value = true
      await store.dispatch(Action.AuthActionTypes.AUTHENTICATE_USER, input)
      loading.value = false
      if (store.getters['isLoggedIn']) {
        if (route.query && route.query.redirectTo) {
          router.push(route.query.redirectTo as string)
        } else {
          router.push('/home')
        }
      }
    }
    return { input, loading, authenticate, inputEmpty }
  },
})
</script>

<template>
  <div
    class="
      flex
      justify-center
      container
      mx-auto
      h-screen
      w-1/4
      px-4
      lg:px-0
      py-8
    "
  >
    <div>
      <IconTwitter :size="60" class="h-12 w-12 text-blue" />
      <h1 class="pt-12 text-4xl dark:text-lightest font-bold">
        Log in to Twitter
      </h1>
      <form @submit.prevent="authenticate" class="w-full text-center">
        <input
          v-model="input.email"
          type="text"
          placeholder="Email"
          class="
            w-full
            px-2
            py-4
            my-6
            border-2 border-lighter
            text-xl
            rounded
            dark:border-dark
            focus:outline-none
            dark:bg-black
            dark:text-light
            focus:border-blue
            dark:focus:border-blue
            transition-colors
            duration-75
          "
        />
        <input
          v-model="input.password"
          type="password"
          placeholder="Password"
          class="
            w-full
            px-2
            py-4
            mb-6
            border-2 border-lighter
            text-xl
            rounded
            dark:border-dark
            focus:outline-none
            dark:bg-black
            dark:text-light
            focus:border-blue
            dark:focus:border-blue
            transition-colors
            duration-75
          "
        />
        <button
          type="submit"
          class="bg-blue rounded-full focus:outline-none w-full h-auto p-4"
          :class="
            inputEmpty
              ? 'cursor-not-allowed'
              : 'cursor-pointer hover:bg-darkblue'
          "
          :disabled="inputEmpty"
        >
          <LoadingSpinner v-if="loading" color="white" size="36px" />
          <span v-else class="text-lightest text-lg font-semibold">Log in</span>
        </button>
        <router-link to="/">
          <span class="text-blue">Sign up for Twitter</span>
        </router-link>
      </form>
    </div>
  </div>
</template>
