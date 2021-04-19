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
import Return from '../../../components/common/Return.vue'
import TweetCard from '../../tweets/components/TweetCard.vue'
import LoadingSpinner from '../../../components/common/LoadingSpinner.vue'
import { ActionTypes } from '../../tweets/store/actions'
import { TweetAndReplies } from '../../tweets/store/state'
import TweetCreateReplyDialog from '../../tweets/components/TweetCreateReplyDialog.vue'

export default defineComponent({
  components: { TweetCard, LoadingSpinner, TweetCreateReplyDialog, Return },
  name: 'Status',
  setup() {
    const store = useStore()
    const route = useRoute()
    const elRef = ref<Element>(null)
    const initialLoadDone = ref<boolean>(false)
    const loadNextBatch = ref<boolean>(false)
    const tweet = ref<TweetAndReplies | null>(null)

    const showCreateReplyDialog = ref<boolean>(false)

    const parsedCreatedAt = computed(() =>
      dayjs(store.getters['tweetStatus'].createdAt).format(
        'h:mm A · MMM D, YYYY'
      )
    )

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

    async function createReply(content: string) {
      try {
        await store.dispatch(ActionTypes.NEW_REPLY, {
          tweetId: route.params.tweetId,
          content,
        })
        initialLoadDone.value = false
        await store.dispatch(ActionTypes.GET_TWEET_STATUS, route.params.tweetId)
        initialLoadDone.value = true

        tweet.value = store.getters['tweetStatus']
      } catch (error) {}
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
      showCreateReplyDialog,
      createReply,
      parsedCreatedAt,
      handleScroll,
    }
  },
})
</script>

<template>
  <TweetCreateReplyDialog
    :show="showCreateReplyDialog"
    @close="showCreateReplyDialog = false"
    @dispatch="createReply"
  />

  <main
    class="w-full h-full overflow-y-scroll border-r border-lighter dark:border-darker md:border-r-0"
    ref="elRef"
    @scroll="handleScroll"
  >
    <div
      class="px-5 py-3 border-b border-lighter dark:border-dark flex items-center justify-start space-x-6"
    >
      <Return />
      <h1 class="text-2xl font-bold dark:text-lightest">Tweet</h1>
    </div>
    <div class="px-5 py-3 border-b border-lighter dark:border-dark">
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
        <p class="text-2xl py-2 break-words dark:text-lightest">
          {{ tweet.content }}
        </p>
        <p class="text-dark dark:text-light">{{ parsedCreatedAt }}</p>
        <div
          class="flex items-center justify-start space-x-12 w-full border-t border-b my-4 py-4 border-lighter dark:border-dark"
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
          class="flex items-center justify-around w-full text-xl text-dark dark:text-light"
        >
          <div
            class="flex justify-center hover:bg-darkblue hover:text-blue hover:bg-opacity-20 rounded-full p-3 cursor-pointer transition duration-75"
            @click="showCreateReplyDialog = true"
          >
            <FontAwesome :icon="['fas', 'comment']" />
          </div>
          <div
            class="flex justify-center hover:bg-success hover:text-success hover:bg-opacity-20 rounded-full p-3 cursor-pointer"
          >
            <FontAwesome :icon="['fas', 'retweet']" />
          </div>
          <div
            class="flex justify-center hover:bg-danger hover:text-danger hover:bg-opacity-20 rounded-full p-3 cursor-pointer"
          >
            <FontAwesome :icon="['fas', 'heart']" />
          </div>
          <div
            class="flex justify-center hover:bg-darkblue hover:text-darkblue hover:bg-opacity-20 rounded-full p-3 cursor-pointer"
          >
            <FontAwesome :icon="['fas', 'share-square']" />
          </div>
        </div>
      </div>
    </div>
    <div v-show="!initialLoadDone" class="flex flex-col">
      <div class="w-full text-center">
        <LoadingSpinner />
      </div>
    </div>
    <div v-if="initialLoadDone && tweet.replies.length > 0">
      <div
        v-for="reply in tweet.replies"
        :key="reply.id"
        class="w-full p-4 border-b dark:border-dark hover:bg-lighter dark:hover:bg-light dark:hover:bg-opacity-20 flex cursor-pointer transition-colors duration-75"
      >
        <div class="w-full">
          <router-link :to="`/${reply.name}/status/${reply.id}`">
            <TweetCard :tweet="reply" />
          </router-link>
        </div>
      </div>
      <div
        v-show="tweet.replies.length > 0 && loadNextBatch"
        class="w-full p-4 border-b dark:border-dark hover:bg-lighter dark:hover:bg-light dark:hover:bg-opacity-20 flex cursor-pointer transition-colors duration-75"
      >
        <div class="w-full text-center">
          <LoadingSpinner />
        </div>
      </div>
    </div>
  </main>
</template>
