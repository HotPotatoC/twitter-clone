<template>
  <div
    class="flex justify-center container mx-auto h-screen w-1/4 px-4 lg:px-0 py-8"
  >
    <div>
      <font-awesome
        :icon="['fab', 'twitter']"
        class="h-12 w-12 text-6xl text-blue rounded-full"
      >
      </font-awesome>
      <h1 class="pt-12 text-4xl dark:text-lightest font-bold">
        Log in to Twitter
      </h1>
      <form @submit.prevent="authenticate" class="w-full">
        <input
          v-model="input.email"
          type="text"
          placeholder="Email"
          class="w-full px-2 py-4 my-6 border-2 border-lighter text-xl rounded dark:border-light dark:border-opacity-25 focus:outline-none dark:bg-black dark:text-light"
        />
        <input
          v-model="input.password"
          type="password"
          placeholder="Password"
          class="w-full px-2 py-4 mb-6 border-2 border-lighter text-xl rounded dark:border-light dark:border-opacity-25 focus:outline-none dark:bg-black dark:text-light"
        />
        <button
          type="submit"
          class="bg-blue text-lightest text-lg rounded-full font-semibold focus:outline-none w-full h-auto p-4"
          :class="
            inputEmpty
              ? 'opacity-50 cursor-default'
              : 'cursor-pointer hover:bg-darkblue'
          "
          :disabled="inputEmpty"
        >
          Log in
        </button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from '../../store'
import { ActionTypes } from './store/actions'

export default defineComponent({
  name: 'Login',
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()

    const email = ref<string>('')
    const password = ref<string>('')

    const input = reactive({ email, password })

    const inputEmpty = computed(
      () => input.email === '' || input.password === ''
    )

    async function authenticate() {
      await store.dispatch(ActionTypes.AUTHENTICATE_USER, input)
      if (store.getters['isLoggedIn']) {
        if (route.query && route.query.redirectTo) {
          router.push(route.query.redirectTo as string)
        } else {
          router.push('/home')
        }
      }
    }
    return { input, authenticate, inputEmpty }
  },
})
</script>
