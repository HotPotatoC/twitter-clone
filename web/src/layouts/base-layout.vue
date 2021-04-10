<template>
  <div class="flex container mx-auto px-12 h-screen w-full font-sans">
    <navigation-sidebar />
    <router-view />
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
