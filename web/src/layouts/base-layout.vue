<template>
  <div class="flex container mx-auto px-12 h-screen w-full font-sans">
    <navigation-sidebar />
    <router-view />
    <trending-sidebar />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import NavigationSidebar from '../components/layout/navigation-sidebar.vue'
import TrendingSidebar from '../components/layout/trending-sidebar.vue'

export default defineComponent({
  components: { NavigationSidebar, TrendingSidebar },
  name: 'BaseLayout',
})
</script>
