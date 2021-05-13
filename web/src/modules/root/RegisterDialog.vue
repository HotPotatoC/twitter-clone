<script lang="ts">
import { defineComponent, ref, reactive, computed } from 'vue'
import Dialog from '../../shared/Dialog.vue'
import IconTwitterWhite from '../../icons/IconTwitterWhite.vue'

export default defineComponent({
  name: 'RegisterDialog',
  components: { Dialog, IconTwitterWhite },
  props: {
    show: {
      type: Boolean,
      default: false,
    },
  },
  emit: ['close', 'dispatch'],
  setup(props, { emit }) {
    const handle = ref('')
    const email = ref('')
    const password = ref('')
    const registerContent = reactive({
      handle,
      email,
      password,
    })

    const contentIsEmpty = computed(
      () => handle.value === '' || email.value === '' || password.value === ''
    )

    function close() {
      emit('close')
    }

    function dispatch() {
      emit('dispatch', registerContent)
    }

    return {
      registerContent,
      contentIsEmpty,
      close,
      dispatch,
    }
  },
})
</script>

<template>
  <Dialog :show="show" @close="close" :closeButton="false" size="xl">
    <template #title>
      <div class="text-center mb-6">
        <IconTwitterWhite :size="36" />
      </div>
      <h1 class="font-bold text-lightest text-2xl mb-6">Create your account</h1>
    </template>

    <form @submit.prevent="dispatch" class="w-full mb-6">
      <input
        v-model="registerContent.handle"
        type="text"
        placeholder="Username"
        class="
          w-full
          px-2
          py-4
          mb-8
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
        v-model="registerContent.email"
        type="email"
        placeholder="Email"
        class="
          w-full
          px-2
          py-4
          mb-8
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
        v-model="registerContent.password"
        type="password"
        placeholder="Password"
        class="
          w-full
          px-2
          py-4
          mb-8
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
        class="
          w-full
          p-4
          text-lg
          rounded-full
          font-semibold
          focus:outline-none
          transition-colors
          duration-75
        "
        :class="
          contentIsEmpty
            ? ['bg-dark', 'text-light', 'cursor-default']
            : ['bg-blue', 'hover:bg-darkblue', 'text-lightest']
        "
        @click="close"
      >
        Sign Up
      </button>
    </form>
  </Dialog>
</template>
