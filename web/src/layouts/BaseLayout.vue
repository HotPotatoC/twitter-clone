<template>
  <div class="flex container mx-auto px-4 xl:px-40 h-screen w-full font-sans">
    <NavigationSidebar />
    <router-view />
    <TrendingSidebar />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import NavigationSidebar from '../components/layout/NavigationSidebar.vue'
import TrendingSidebar from '../components/layout/TrendingSidebar.vue'

export default defineComponent({
  components: { NavigationSidebar, TrendingSidebar },
  name: 'BaseLayout',
})
</script>
