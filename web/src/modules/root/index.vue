<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from '../../store'
import { Action } from '../storeActionTypes'
import RegisterDialog from './RegisterDialog.vue'
import IconTwitterWhite from '../../icons/IconTwitterWhite.vue'

interface RegisterData {
  handle: string
  email: string
  password: string
}

export default defineComponent({
  name: 'Root',
  components: { RegisterDialog, IconTwitterWhite },
  setup() {
    const store = useStore()
    const router = useRouter()
    const showRegisterDialog = ref(false)

    async function register({ handle, email, password }: RegisterData) {
      await store.dispatch(Action.UserActionTypes.REGISTER_ACCOUNT, {
        handle,
        email,
        password,
      })
      router.push('/home')
    }

    return { showRegisterDialog, register }
  },
})
</script>

<template>
  <RegisterDialog
    :show="showRegisterDialog"
    @close="showRegisterDialog = false"
    @dispatch="register"
  />
  <div class="flex flex-wrap flex-col-reverse md:flex-row">
    <div
      class="w-full md:w-1/2 h-full relative flex justify-center items-center"
    >
      <img
        src="../../assets/images/twitter-banner.png"
        class="w-full h-screen object-cover select-none"
        alt=""
      />
      <IconTwitterWhite :size="326" class="absolute" />
    </div>
    <div class="w-full md:w-1/2 h-screen px-6 md:px-12">
      <IconTwitterWhite :size="56" class="mt-32" />
      <h1 class="mt-32 dark:text-lighter text-5xl md:text-7xl font-bold">
        Happening now
      </h1>
      <h2 class="mt-12 dark:text-lighter text-2xl md:text-4xl font-bold">
        Join Twitter today.
      </h2>

      <div class="flex flex-wrap flex-col space-y-6 mt-12">
        <button
          class="
            bg-blue
            text-lightest text-lg
            rounded-full
            font-semibold
            focus:outline-none
            h-auto
            p-4
            hover:bg-darkblue
            transition-colors
            duration-75
          "
          @click="showRegisterDialog = true"
        >
          Sign Up
        </button>
        <router-link
          to="/login"
          class="
            border-2 border-blue
            text-blue
            text-lg
            rounded-full
            font-semibold
            focus:outline-none
            h-auto
            p-4
            hover:bg-blue hover:bg-opacity-10
            transition-colors
            duration-75
            text-center
          "
        >
          Log in
        </router-link>
      </div>
    </div>
  </div>
</template>
