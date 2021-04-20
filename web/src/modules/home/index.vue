<script lang="ts">
import {
  computed,
  defineComponent,
  onBeforeMount,
  reactive,
  Ref,
  ref,
} from 'vue'
import { useStore } from '../../store'
import { Tweet } from '../tweets/store/state'
import { ActionTypes } from '../tweets/store/actions'
import TweetCard from '../tweets/components/TweetCard.vue'
import LoadingSpinner from '../../components/common/LoadingSpinner.vue'
import IconStar from '../../components/icons/IconStar.vue'

interface NewTweet {
  content: string | Ref<string>
}

export default defineComponent({
  components: { TweetCard, LoadingSpinner, IconStar },
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

    const newTweetContentIsEmpty = computed(() => tweetContent.value === '')

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
      newTweetContentIsEmpty,
      tweetsRef,
      addNewTweet,
      handleScroll,
    }
  },
})
</script>

<template>
  <main
    class="w-full h-full overflow-y-scroll border-r border-lighter dark:border-darker md:border-r-0"
    ref="tweetsRef"
    @scroll="handleScroll"
  >
    <div
      class="px-5 py-3 border-b border-lighter dark:border-dark flex items-center justify-between"
    >
      <h1 class="text-xl font-bold dark:text-lightest">Home</h1>
      <IconStar :size="20" class="text-blue fill-current" />
    </div>
    <div class="px-5 py-3 border-b-8 border-lighter dark:border-dark flex">
      <form @submit.prevent="addNewTweet" class="w-full px-4 relative">
        <textarea
          v-model="newTweet.content"
          placeholder="What's happening?"
          class="mt-3 pb-3 w-full focus:outline-none dark:bg-black dark:text-light"
        />
        <button
          type="submit"
          class="h-10 px-4 font-semibold focus:outline-none rounded-full absolute bottom-0 right-0 transition-colors duration-75"
          :class="
            newTweetContentIsEmpty
              ? ['bg-dark', 'text-light', 'cursor-default']
              : ['bg-blue', 'hover:bg-darkblue', 'text-lightest']
          "
          :disabled="newTweetContentIsEmpty"
        >
          Tweet
        </button>
      </form>
    </div>
    <div v-show="!initialLoadDone" class="flex flex-col">
      <div class="w-full text-center">
        <LoadingSpinner />
      </div>
    </div>
    <div v-show="initialLoadDone" class="flex flex-col">
      <div
        v-for="tweet in tweets"
        :key="tweet.id"
        class="w-full p-4 border-b dark:border-dark hover:bg-lighter dark:hover:bg-darker flex cursor-pointer transition-colors duration-75"
      >
        <div class="w-full">
          <router-link :to="`/${tweet.handle}/status/${tweet.id}`">
            <TweetCard :tweet="tweet" />
          </router-link>
        </div>
      </div>

      <div
        v-show="tweets.length > 0 && loadNextBatch"
        class="w-full p-4 border-b dark:border-dark hover:bg-lighter dark:hover:bg-darker flex cursor-pointer"
      >
        <div class="w-full text-center">
          <LoadingSpinner />
        </div>
      </div>
    </div>
  </main>
</template>
