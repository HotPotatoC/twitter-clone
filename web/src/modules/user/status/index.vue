<template>
  <main
    class="w-full h-full overflow-y-scroll border-r border-lighter dark:border-light dark:border-opacity-25"
    ref="elRef"
    @scroll="handleScroll"
  >
    <div
      class="px-5 py-3 border-b border-lighter dark:border-light dark:border-opacity-25 flex items-center justify-start space-x-6"
    >
      <FontAwesome :icon="['fas', 'arrow-left']" class="text-xl text-blue" />
      <h1 class="text-2xl font-bold dark:text-lightest">Tweet</h1>
    </div>
    <div
      class="px-5 py-3 border-b border-lighter dark:border-light dark:border-opacity-25"
    >
      <div v-if="initialLoadDone" class="w-full">
        <div class="flex items-center w-full">
          <div class="block">
            <p class="font-semibold dark:text-lightest">{{ tweet.name }}</p>
            <p class="text-sm text-dark dark:text-light">@{{ tweet.name }}</p>
          </div>
          <FontAwesome
            :icon="['fas', 'angle-down']"
            class="text-dark ml-auto"
          />
        </div>
        <p class="text-2xl py-2 dark:text-lightest">
          {{ tweet.content }}
        </p>
        <p class="text-dark dark:text-light">{{ parsedCreatedAt }}</p>
        <div
          class="flex items-center justify-start space-x-12 w-full border-t border-b my-4 py-4 border-lighter dark:border-light dark:border-opacity-25"
        >
          <div class="flex space-x-2 text-sm">
            <p class="text-light dark:text-lightest font-bold">
              {{ tweet.repliesCount }}
            </p>
            <p class="text-dark dark:text-light">Retweets</p>
          </div>
          <div class="flex space-x-2 text-sm">
            <p class="text-light dark:text-lightest font-bold">
              {{ tweet.favoritesCount }}
            </p>
            <p class="text-dark dark:text-light">Likes</p>
          </div>
        </div>
        <div
          class="flex items-center justify-around w-full mb-2 text-xl text-dark dark:text-light"
        >
          <FontAwesome :icon="['fas', 'comment']" class="mr-3" />
          <FontAwesome :icon="['fas', 'retweet']" class="mr-3" />
          <FontAwesome :icon="['fas', 'heart']" class="mr-3" />
          <FontAwesome :icon="['fas', 'share-square']" class="mr-3" />
        </div>
      </div>
    </div>
    <div v-if="initialLoadDone && tweet.replies.length > 0">
      <div
        v-for="reply in tweet.replies"
        :key="reply.id"
        class="w-full p-4 border-b dark:border-light dark:border-opacity-25 hover:bg-lighter dark:hover:bg-light dark:hover:bg-opacity-20 flex cursor-pointer"
      >
        <div class="w-full">
          <router-link :to="`/${reply.name}/status/${reply.id}`">
            <tweet-card :tweet="reply" />
          </router-link>
        </div>
      </div>
      <div
        v-show="tweet.replies.length > 0 && loadNextBatch"
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
import dayjs from 'dayjs'
import {
  computed,
  defineComponent,
  onBeforeMount,
  onMounted,
  ref,
  watch,
} from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from '../../../store'
import TweetCard from '../../../components/common/TweetCard.vue'
import LoadingSpinner from '../../../components/common/LoadingSpinner.vue'
import { ActionTypes } from '../../tweets/store/actions'
import { TweetAndReplies } from '../../tweets/store/state'

export default defineComponent({
  components: { TweetCard, LoadingSpinner },
  name: 'Status',
  setup() {
    const store = useStore()
    const route = useRoute()
    const elRef = ref<Element>(null)
    const initialLoadDone = ref<boolean>(false)
    const loadNextBatch = ref<boolean>(false)
    const tweet = ref<TweetAndReplies | null>(null)

    onBeforeMount(async () => {
      await getTweetStatus(route.params.tweetId)
      initialLoadDone.value = true
    })

    onMounted(async () => {
      watch(
        () => route.params.tweetId,
        async (tweetId) => {
          initialLoadDone.value = false
          await getTweetStatus(tweetId)
          initialLoadDone.value = true
        },
        { flush: 'post' }
      )
    })

    const parsedCreatedAt = computed(() =>
      dayjs(store.getters['tweetStatus'].createdAt).format(
        'h:mm A · MMM D, YYYY'
      )
    )

    async function getTweetStatus(tweetId: string | string[]) {
      await store.dispatch(ActionTypes.GET_TWEET_STATUS, tweetId)

      tweet.value = store.getters['tweetStatus']
    }

    async function loadMoreReplies(tweetId: string | string[]) {
      if (
        initialLoadDone.value &&
        store.getters['tweetStatus'].replies.length > 0
      ) {
        const cursor = store.getters['lastStatusReplyItem'].createdAt
        await store.dispatch(ActionTypes.LOAD_MORE_REPLIES, { tweetId, cursor })
      }
    }

    async function handleScroll(e: Event) {
      const element = elRef.value
      if (
        !loadNextBatch.value &&
        element.scrollTop + element.clientHeight + 1 >= element.scrollHeight
      ) {
        loadNextBatch.value = true
        await loadMoreReplies(route.params.tweetId)
        loadNextBatch.value = false
      }
    }

    return {
      elRef,
      initialLoadDone,
      loadNextBatch,
      tweet,
      parsedCreatedAt,
      handleScroll,
    }
  },
})
</script>
