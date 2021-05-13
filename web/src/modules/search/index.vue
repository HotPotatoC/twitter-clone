<script lang="ts">
import { defineComponent, onBeforeMount, onMounted, Ref, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useStore } from '../../store'
import { Tweet } from '../tweets/types'
import { Action } from '../storeActionTypes'
import TweetCard from '../tweets/TweetCard.vue'
import LoadingSpinner from '../../shared/LoadingSpinner.vue'
import Return from '../../shared/Return.vue'
import IconSearch from '../../icons/IconSearch.vue'

interface NewTweet {
  content: string | Ref<string>
}

export default defineComponent({
  components: { TweetCard, LoadingSpinner, Return, IconSearch },
  name: 'Home',
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()
    const initialLoadDone = ref(false)
    const loadNextBatch = ref(false)
    const tweetsRef = ref<Element>()
    const tweets = ref<Tweet[]>([])
    const searchFocused = ref(false)
    const searchQuery = ref('')

    onBeforeMount(async () => {
      await loadSearchResults()
      initialLoadDone.value = true
    })

    onMounted(() => {
      watch(
        () => route.query.q,
        async () => {
          initialLoadDone.value = false
          await loadSearchResults()
          initialLoadDone.value = true
        },
        { flush: 'post' }
      )
    })

    function redirectWithSearchQuery() {
      router.push({
        path: '/search',
        query: { q: searchQuery.value },
      })
      return
    }

    async function loadSearchResults() {
      // if (initialLoadDone.value && store.getters['TweetsFeed'].length > 0) {
      //   const lastItem = store.getters['LastTweetFeedItem']
      //   await store.dispatch(ActionTypes.LOAD_MORE_TWEETS, lastItem.createdAt)
      // } else {
      await store.dispatch(
        Action.TweetsActionTypes.SEARCH_TWEETS,
        route.query.q as string
      )
      // }

      tweets.value = store.getters['tweetsSearchResults']
    }

    async function handleScroll(e: Event) {
      // const element = tweetsRef.value
      // if (
      //   !loadNextBatch.value &&
      //   element.scrollTop + element.clientHeight + 1 >= element.scrollHeight
      // ) {
      //   loadNextBatch.value = true
      //   await loadSearchResults()
      //   loadNextBatch.value = false
      // }
    }

    return {
      searchFocused,
      searchQuery,
      redirectWithSearchQuery,
      initialLoadDone,
      loadNextBatch,
      tweets,
      tweetsRef,
      handleScroll,
    }
  },
})
</script>

<template>
  <main
    class="
      w-full
      h-full
      overflow-y-scroll
      border-r border-lighter
      dark:border-darker
      md:border-r-0
    "
    ref="tweetsRef"
    @scroll="handleScroll"
  >
    <div
      class="
        relative
        px-5
        py-3
        border-b border-lighter
        dark:border-dark
        flex
        items-center
        justify-between
      "
    >
      <Return class="mr-6" />
      <form @submit.prevent="redirectWithSearchQuery" class="w-full">
        <input
          class="
            pl-12
            rounded-full
            w-full
            p-2
            bg-lighter
            dark:bg-darkest
            dark:text-light
            text-sm
            focus:bg-white
            dark:focus:bg-black
            focus:outline-none
            border-2 border-lighter
            dark:border-darkest
            focus:border-blue
            dark:focus:border-blue
            dark:focus:text-lightest
            transition
            duration-150
          "
          @focus="searchFocused = true"
          @blur="searchFocused = false"
          v-model="searchQuery"
          type="search"
          placeholder="Search Twitter"
        />
        <input type="submit" class="hidden" />
      </form>
      <IconSearch
        :size="24"
        class="absolute left-0 top-0 mt-5 ml-20"
        :class="searchFocused ? 'text-blue' : 'text-light'"
      />
    </div>
    <div v-show="initialLoadDone" class="flex flex-col">
      <TweetCard :tweet="tweet" v-for="tweet in tweets" :key="tweet.id" />

      <!-- <div
        v-show="tweets.length > 0 && loadNextBatch"
        class="w-full p-4 border-b dark:border-dark hover:bg-lighter dark:hover:bg-light dark:hover:bg-opacity-20 flex cursor-pointer"
      >
        <div class="w-full text-center">
          <loading-spinner />
        </div>
      </div> -->
    </div>
  </main>
</template>
