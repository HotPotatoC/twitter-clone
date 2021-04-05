<template>
  <div
    class="px-5 py-3 border-b border-lighter dark:border-light dark:border-opacity-25 flex items-center justify-start space-x-6"
  >
    <font-awesome :icon="['fas', 'arrow-left']" class="text-xl text-blue" />
    <h1 class="text-2xl font-bold dark:text-white">Tweet</h1>
  </div>
  <div
    class="px-5 py-3 border-b border-lighter dark:border-light dark:border-opacity-25"
  >
    <div v-if="ready" class="w-full">
      <div class="flex items-center w-full">
        <div class="block">
          <p class="font-semibold dark:text-lightest">{{ tweet.name }}</p>
          <p class="text-sm text-dark dark:text-light">@{{ tweet.name }}</p>
        </div>
        <font-awesome :icon="['fas', 'angle-down']" class="text-dark ml-auto" />
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
      <div class="flex items-center justify-around w-full mb-2">
        <div class="flex items-center text-xl text-dark dark:text-light">
          <font-awesome :icon="['fas', 'comment']" class="mr-3" />
        </div>
        <div class="flex items-center text-xl text-dark dark:text-light">
          <font-awesome :icon="['fas', 'retweet']" class="mr-3" />
        </div>
        <div class="flex items-center text-xl text-dark dark:text-light">
          <font-awesome :icon="['fas', 'heart']" class="mr-3" />
        </div>
        <div class="flex items-center text-xl text-dark dark:text-light">
          <font-awesome :icon="['fas', 'share-square']" class="mr-3" />
        </div>
      </div>
    </div>
  </div>
  <div v-if="ready && tweet.replies.length > 0">
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
  </div>
</template>

<script lang="ts">
import dayjs from 'dayjs'
import { computed, defineComponent, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from '../../../store'
import { ActionTypes } from '../../../store/tweets/actions'
import TweetCard from '../../../components/common/TweetCard.vue'

export default defineComponent({
  components: { TweetCard },
  name: 'Status',
  setup() {
    const store = useStore()
    const route = useRoute()
    const ready = ref<boolean>(false)

    onMounted(async () => {
      await store.dispatch(ActionTypes.GET_TWEET_STATUS, route.params.tweetId)

      ready.value = true
    })

    const tweet = computed(() => store.getters['getTweetStatus'])
    const parsedCreatedAt = computed(() =>
      dayjs(store.getters['getTweetStatus'].createdAt).format(
        'h:mm A · MMM D, YYYY'
      )
    )

    return { ready, tweet, parsedCreatedAt }
  },
})
</script>
