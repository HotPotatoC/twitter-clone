<script lang="ts">
import {
  defineComponent,
  onBeforeMount,
  onMounted,
  ref,
  watchEffect,
} from 'vue'
import { useStore } from '../../store'
import { useScroll } from '../../hooks/useScroll'
import { Tweet } from '../tweets/types'
import { Action } from '../storeActionTypes'
import TweetCard from '../tweets/TweetCard.vue'
import TweetConversationCard from '../tweets/TweetConversationCard.vue'
import LoadingSpinner from '../../shared/LoadingSpinner.vue'
import IconStar from '../../icons/IconStar.vue'
import TweetCreateTweetForm from '../tweets/TweetCreateTweetForm.vue'

export default defineComponent({
  components: {
    TweetCard,
    TweetConversationCard,
    LoadingSpinner,
    IconStar,
    TweetCreateTweetForm,
  },
  name: 'Home',
  setup() {
    const store = useStore()
    const initialLoadDone = ref(false)
    const loadNextBatch = ref(false)
    const tweets = ref<Tweet[]>([])

    const [scrollRef, isBottom] = useScroll()

    onBeforeMount(async () => {
      await loadTweets()
      initialLoadDone.value = true
    })

    onMounted(() => {
      watchEffect(async () => {
        if (!loadNextBatch.value && isBottom.value) {
          loadNextBatch.value = true
          await loadTweets()
          loadNextBatch.value = false
          isBottom.value = false
        }
      })
    })

    async function loadTweets() {
      if (initialLoadDone.value && store.getters['tweetsFeed'].length > 0) {
        const lastItem = store.getters['lastTweetFeedItem']
        await store.dispatch(
          Action.TweetsActionTypes.LOAD_MORE_TWEETS,
          lastItem.createdAt
        )
      } else {
        await store.dispatch(Action.TweetsActionTypes.GET_TWEETS_FEED)
      }

      tweets.value = store.getters['tweetsFeed']
    }

    async function addNewTweet({ content, attachments }) {
      try {
        initialLoadDone.value = false
        await store.dispatch(Action.TweetsActionTypes.NEW_TWEET, {
          content,
          attachments,
        })
        await store.dispatch(Action.TweetsActionTypes.GET_TWEETS_FEED)
        initialLoadDone.value = true
      } catch (error) {
        console.log(error)
      }
    }

    return {
      initialLoadDone,
      loadNextBatch,
      tweets,
      scrollRef,
      addNewTweet,
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
    ref="scrollRef"
  >
    <div
      class="
        px-5
        py-3
        border-b border-lighter
        dark:border-dark
        flex
        items-center
        justify-between
      "
    >
      <h1 class="text-xl font-bold dark:text-lightest">Home</h1>
      <IconStar :size="20" class="text-blue fill-current" />
    </div>
    <div class="px-5 py-3 border-b-8 border-lighter dark:border-dark flex">
      <TweetCreateTweetForm @submit="addNewTweet" />
    </div>
    <div v-show="!initialLoadDone" class="flex flex-col">
      <div class="w-full text-center">
        <LoadingSpinner />
      </div>
    </div>
    <div v-show="initialLoadDone" class="flex flex-col">
      <div v-for="tweet in tweets" :key="tweet.id">
        <TweetCard v-if="!tweet.isReply" :tweet="tweet" />
        <TweetConversationCard v-else :tweet="tweet" />
      </div>

      <div
        v-show="tweets.length > 0 && loadNextBatch"
        class="
          w-full
          p-4
          border-b border-lighter
          dark:border-dark
          hover:bg-lighter
          dark:hover:bg-darker
          cursor-pointer
        "
      >
        <div class="w-full text-center">
          <LoadingSpinner />
        </div>
      </div>
    </div>
  </main>
</template>
