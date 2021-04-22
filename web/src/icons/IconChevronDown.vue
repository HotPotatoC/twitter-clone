<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'IconArrowDown',
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
    class="feather feather-chevron-down"
  >
    <polyline points="6 9 12 15 18 9"></polyline>
  </svg>
</template>
