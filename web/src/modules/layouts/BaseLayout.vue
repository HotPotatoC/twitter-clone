<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '../../store'
import { Action } from '../storeActionTypes'
import TweetImageOverlay from '../tweets/TweetImageOverlay.vue'
import NavigationSidebar from './NavigationSidebar.vue'
import TrendingSidebar from './TrendingSidebar.vue'

export default defineComponent({
  components: { NavigationSidebar, TrendingSidebar, TweetImageOverlay },
  name: 'BaseLayout',
  setup() {
    const store = useStore()

    const tweetImageOverlay = computed(() => store.getters['tweetImageOverlay'])

    function closeTweetImageOverlay() {
      store.dispatch(Action.TweetsActionTypes.TOGGLE_TWEET_IMAGE_OVERLAY, {
        tweet: null,
        show: false,
        source: '',
      })
    }

    return { tweetImageOverlay, closeTweetImageOverlay }
  },
})
</script>

<template>
  <TweetImageOverlay
    :image="tweetImageOverlay.source"
    :show="tweetImageOverlay.show"
    :tweet="tweetImageOverlay.tweet"
    @close="closeTweetImageOverlay"
  />
  <div class="flex container mx-auto px-4 xl:px-40 h-screen w-full font-sans">
    <NavigationSidebar />
    <router-view />
    <TrendingSidebar />
  </div>
</template>
