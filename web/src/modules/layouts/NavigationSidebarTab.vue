<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'NavigationSidebarTab',
  props: {
    selected: {
      type: Boolean,
      default: false,
    },
    id: {
      type: String,
      required: true,
    },
    label: {
      type: String,
      required: true,
    },
    to: {
      type: String,
      required: true,
    },
  },
  emits: ['selectTab'],
  setup(_, { emit }) {
    return { emit }
  },
})
</script>

<template>
  <router-link :to="to">
    <button
      @click="emit('selectTab', id)"
      class="
        flex
        items-center
        mr-auto
        mb-3
        py-2
        md:px-4
        rounded-full
        hover:bg-lightblue
        dark:hover:bg-darkblue dark:hover:bg-opacity-20
        hover:text-blue
        focus:outline-none
        transition-colors
        duration-75
      "
      :class="
        selected ? ['text-blue', 'dark:text-blue'] : ['dark:text-lighter']
      "
    >
      <slot name="icon"></slot>
      <span class="text-left text-lg font-bold hidden lg:block">
        {{ label }}
      </span>
    </button>
  </router-link>
</template>
