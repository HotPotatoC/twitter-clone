<template>
  <div class="flex container mx-auto px-12 h-screen w-full font-sans">
    <navigation-sidebar />
    <main
      class="w-full h-full overflow-y-scroll hide-scrollbar border-r border-lighter dark:border-light dark:border-opacity-25"
    >
      <router-view />
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import NavigationSidebar from '../components/layout/navigation-sidebar.vue'

export default defineComponent({
  components: { NavigationSidebar },
  name: 'BaseLayout',
})
</script>
