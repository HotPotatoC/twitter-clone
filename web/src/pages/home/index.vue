<template>
  <div
    class="px-5 py-3 border-b border-lighter dark:border-light dark:border-opacity-25 flex items-center justify-between"
  >
    <h1 class="text-xl font-bold dark:text-lightest">Home</h1>
    <font-awesome :icon="['fas', 'star']" class="text-xl text-blue" />
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
  <div v-if="ready" class="flex flex-col-reverse">
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
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, reactive, Ref, ref } from 'vue'
import axios from '../../services/axios'
import { useStore } from '../../store'
import { ActionTypes } from '../../store/tweets/actions'
import TweetCard from '../../components/common/TweetCard.vue'

interface NewTweet {
  content: string | Ref<string>
}

export default defineComponent({
  components: { TweetCard },
  name: 'Home',
  setup() {
    const store = useStore()
    const ready = ref<boolean>(false)
    const tweetContent = ref<string>('')
    const newTweet = reactive<NewTweet>({
      content: tweetContent,
    })

    async function addNewTweet() {
      await axios.post('/tweets', {
        content: newTweet.content,
      })
    }

    onMounted(async () => {
      await store.dispatch(ActionTypes.GET_TWEETS_FEED)
      ready.value = true
    })

    const tweets = computed(() =>
      store.getters['getTweetsFeed'].sort(
        (a, b) =>
          new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime()
      )
    )

    return { ready, tweets, newTweet, addNewTweet }
  },
})
</script>

<style></style>
