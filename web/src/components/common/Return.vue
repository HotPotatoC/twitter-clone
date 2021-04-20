<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'
import IconArrowLeft from '../icons/IconArrowLeft.vue'

export default defineComponent({
  components: { IconArrowLeft },
  name: 'Return',
  setup() {
    const { back } = useRouter()

    return { back }
  },
})
</script>

<template>
  <button class="focus:outline-none" @click="back">
    <IconArrowLeft :size="24" class="text-blue" />
  </button>
</template>
