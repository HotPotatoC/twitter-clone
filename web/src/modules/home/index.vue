<template>
  <main
    class="w-full h-full overflow-y-scroll border-r border-lighter dark:border-light dark:border-opacity-25"
    ref="tweetsRef"
    @scroll="handleScroll"
  >
    <div
      class="px-5 py-3 border-b border-lighter dark:border-light dark:border-opacity-25 flex items-center justify-between"
    >
      <h1 class="text-xl font-bold dark:text-lightest">Home</h1>
      <FontAwesome :icon="['fas', 'star']" class="text-xl text-blue" />
    </div>
    <div
      class="px-5 py-3 border-b-8 border-lighter dark:border-light dark:border-opacity-25 flex"
    >
      <form @submit.prevent="addNewTweet" class="w-full px-4 relative">
        <textarea
          v-model="newTweet.content"
          placeholder="What's happening?"
          class="mt-3 pb-3 w-full focus:outline-none dark:bg-black dark:text-light"
        />
        <button
          type="submit"
          class="h-10 px-4 text-lightest font-semibold bg-blue hover:bg-darkblue focus:outline-none rounded-full absolute bottom-0 right-0"
        >
          Tweet
        </button>
      </form>
    </div>
    <div v-show="initialLoadDone" class="flex flex-col">
      <div
        v-for="tweet in tweets"
        :key="tweet.id"
        class="w-full p-4 border-b dark:border-light dark:border-opacity-25 hover:bg-lighter dark:hover:bg-light dark:hover:bg-opacity-20 flex cursor-pointer"
      >
        <div class="w-full">
          <router-link :to="`/${tweet.name}/status/${tweet.id}`">
            <tweet-card :tweet="tweet" />
          </router-link>
        </div>
      </div>

      <div
        v-show="tweets.length > 0 && loadNextBatch"
        class="w-full p-4 border-b dark:border-light dark:border-opacity-25 hover:bg-lighter dark:hover:bg-light dark:hover:bg-opacity-20 flex cursor-pointer"
      >
        <div class="w-full text-center">
          <loading-spinner />
        </div>
      </div>
    </div>
  </main>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, reactive, Ref, ref } from 'vue'
import { useStore } from '../../store'
import TweetCard from '../../components/common/TweetCard.vue'
import LoadingSpinner from '../../components/common/LoadingSpinner.vue'
import { Tweet } from '../tweets/store/state'
import { ActionTypes } from '../tweets/store/actions'

interface NewTweet {
  content: string | Ref<string>
}

export default defineComponent({
  components: { TweetCard, LoadingSpinner },
  name: 'Home',
  setup() {
    const store = useStore()
    const initialLoadDone = ref<boolean>(false)
    const loadNextBatch = ref<boolean>(false)
    const tweetsRef = ref<Element>(null)
    const tweets = ref<Tweet[]>([])

    const tweetContent = ref<string>('')
    const newTweet = reactive<NewTweet>({
      content: tweetContent,
    })

    onBeforeMount(async () => {
      await loadTweets()
      initialLoadDone.value = true
    })

    async function loadTweets() {
      if (initialLoadDone.value && store.getters['tweetsFeed'].length > 0) {
        const lastItem = store.getters['lastTweetFeedItem']
        await store.dispatch(ActionTypes.LOAD_MORE_TWEETS, lastItem.createdAt)
      } else {
        await store.dispatch(ActionTypes.GET_TWEETS_FEED)
      }

      tweets.value = store.getters['tweetsFeed']
    }

    async function addNewTweet() {
      try {
        initialLoadDone.value = false
        await store.dispatch(ActionTypes.NEW_TWEET, newTweet.content)
        await store.dispatch(ActionTypes.GET_TWEETS_FEED)
        initialLoadDone.value = true
        newTweet.content = ''
      } catch (error) {
        console.log(error)
      }
    }

    async function handleScroll(e: Event) {
      const element = tweetsRef.value
      if (
        !loadNextBatch.value &&
        element.scrollTop + element.clientHeight + 1 >= element.scrollHeight
      ) {
        loadNextBatch.value = true
        await loadTweets()
        loadNextBatch.value = false
      }
    }

    return {
      initialLoadDone,
      loadNextBatch,
      tweets,
      newTweet,
      tweetsRef,
      addNewTweet,
      handleScroll,
    }
  },
})
</script>

<style></style>
