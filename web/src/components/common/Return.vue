<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'Return',
  setup() {
    const { back } = useRouter()

    return { back }
  },
})
</script>

<template>
  <button class="focus:outline-none" @click="back">
    <FontAwesome :icon="['fas', 'arrow-left']" class="text-xl text-blue" />
  </button>
</template>
