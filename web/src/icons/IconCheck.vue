<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'IconCheck',
  props: {
    size: {
      type: Number,
      default: 18,
    },
  },
})
</script>

<template>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    :width="size"
    :height="size"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-check"
  >
    <polyline points="20 6 9 17 4 12"></polyline>
  </svg>
</template>
