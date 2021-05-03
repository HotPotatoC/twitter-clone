<script lang="ts">
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import { ref, computed, defineComponent, toRefs } from 'vue'
import { useStore } from '../../store'
import { Tweet } from './types'
import IconEllipsisH from '../../icons/IconEllipsisH.vue'
import IconComment from '../../icons/IconComment.vue'
import IconRetweet from '../../icons/IconRetweet.vue'
import IconHeart from '../../icons/IconHeart.vue'
import IconShare from '../../icons/IconShare.vue'
import { Action } from '../storeActionTypes'
import { linkifyHTMLText } from '../../utils/linkify'
import { useRouter } from 'vue-router'
import TweetImageOverlay from './TweetImageOverlay.vue'

export default defineComponent({
  name: 'TweetCard',
  components: {
    IconComment,
    IconEllipsisH,
    IconRetweet,
    IconHeart,
    IconShare,
    TweetImageOverlay,
  },
  props: {
    tweet: {
      type: Object as () => Tweet,
      required: true,
    },
  },
  setup(props) {
    const store = useStore()
    const router = useRouter()
    const { tweet } = toRefs(props)
    const favoritesCount = ref(tweet.value.favoritesCount)
    const alreadyLiked = ref(tweet.value.alreadyLiked)
    const showTweetImageOverlay = ref(false)
    const selectedImageURL = ref('')

    const parsedCreatedAt = computed(() => {
      dayjs.extend(relativeTime)
      return dayjs(tweet.value.createdAt).fromNow()
    })

    const parsedContent = computed(() => linkifyHTMLText(tweet.value.content))

    async function likeTweet() {
      await store.dispatch(
        Action.TweetsActionTypes.FAVORITE_TWEET,
        tweet.value.id.toString()
      )

      alreadyLiked.value = !alreadyLiked.value
      if (alreadyLiked.value) {
        favoritesCount.value++
      } else {
        favoritesCount.value--
      }
    }

    function showOverlay(imageURL: string) {
      showTweetImageOverlay.value = true
      selectedImageURL.value = imageURL
    }

    return {
      router,
      alreadyLiked,
      favoritesCount,
      parsedCreatedAt,
      parsedContent,
      likeTweet,
      showTweetImageOverlay,
      selectedImageURL,
      showOverlay,
    }
  },
})
</script>

<template>
  <TweetImageOverlay
    :image="selectedImageURL"
    :show="showTweetImageOverlay"
    @close="showTweetImageOverlay = false"
  />
  <div
    class="w-full p-4 border-b border-lighter dark:border-dark hover:bg-lighter dark:hover:bg-darker dark:hover:bg-opacity-30 flex cursor-pointer transition-colors duration-75"
    @click="router.push(`/${tweet.authorHandle}/status/${tweet.id}`)"
  >
    <router-link :to="`/${tweet.authorHandle}`" class="flex-none mr-4">
      <img
        :src="tweet.authorPhotoURL"
        class="h-12 w-12 rounded-full flex-none"
      />
    </router-link>
    <div class="w-full">
      <div class="flex flex-wrap items-center w-full">
        <router-link
          :to="`/${tweet.authorHandle}`"
          class="flex flex-wrap items-center"
        >
          <p class="font-semibold dark:text-lightest hover:underline">
            {{ tweet.authorName }}
          </p>
          <p class="text-sm text-dark dark:text-light ml-2">
            @{{ tweet.authorHandle }} ·
          </p>
          <p class="text-sm text-dark dark:text-light ml-2">
            {{ parsedCreatedAt }}
          </p>
        </router-link>
        <div
          class="text-gray ml-auto p-2 hover:bg-darkblue hover:text-blue hover:bg-opacity-20 rounded-full"
        >
          <IconEllipsisH />
        </div>
      </div>
      <div
        class="py-2 break-words dark:text-lightest"
        v-html="parsedContent"
      ></div>
      <div
        v-if="tweet.photoURLs !== null && tweet.photoURLs.length > 0"
        class="relative overflow-hidden w-full h-96 rounded-lg"
      >
        <div class="box-border relative">
          <div class="grid grid-cols-2 gap-1 h-full">
            <div
              class="w-full"
              :class="tweet.photoURLs.length > 2 ? 'h-48' : 'h-96'"
            >
              <img
                :src="tweet.photoURLs[0]"
                @click.stop="showOverlay(tweet.photoURLs[0])"
                class="object-cover w-full h-full"
              />
              <img
                v-if="tweet.photoURLs.length > 2"
                :src="tweet.photoURLs[1]"
                @click.stop="showOverlay(tweet.photoURLs[1])"
                class="object-cover w-full h-full"
              />
            </div>
            <div
              class="w-full"
              :class="tweet.photoURLs.length > 2 ? 'h-48' : 'h-96'"
            >
              <img
                v-if="tweet.photoURLs.length > 2"
                :src="tweet.photoURLs[2]"
                @click.stop="showOverlay(tweet.photoURLs[2])"
                class="object-cover w-full"
                :class="tweet.photoURLs.length === 4 ? 'h-full' : 'h-96'"
              />
              <img
                v-else
                :src="tweet.photoURLs[1]"
                @click.stop="showOverlay(tweet.photoURLs[1])"
                class="object-cover w-full h-full"
              />
              <img
                v-if="tweet.photoURLs.length === 4"
                :src="tweet.photoURLs[3]"
                @click.stop="showOverlay(tweet.photoURLs[3])"
                class="object-cover w-full h-full"
              />
            </div>
          </div>
        </div>
      </div>
      <div class="flex items-center justify-between w-full mt-2">
        <div
          class="flex items-center group text-dark dark:text-light hover:text-blue dark:hover:text-blue"
        >
          <div
            class="mr-3 p-2 group-hover:bg-darkblue group-hover:bg-opacity-20 rounded-full"
          >
            <IconComment />
          </div>
          <p class="text-sm">
            {{ tweet.repliesCount }}
          </p>
        </div>
        <div
          class="flex items-center group text-dark dark:text-light hover:text-success dark:hover:text-success"
        >
          <div
            class="mr-3 p-2 group-hover:bg-success group-hover:bg-opacity-20 rounded-full"
          >
            <IconRetweet />
          </div>
          <p class="text-sm">
            {{ tweet.repliesCount }}
          </p>
        </div>
        <div
          class="flex items-center group"
          :class="
            alreadyLiked
              ? ['text-danger']
              : [
                  'text-dark',
                  'dark:text-light',
                  'hover:text-danger',
                  'dark:hover:text-danger',
                ]
          "
        >
          <div
            class="mr-3 p-2 group-hover:bg-danger group-hover:bg-opacity-20 rounded-full"
            @click.stop="likeTweet"
          >
            <IconHeart :class="alreadyLiked ? 'fill-current' : null" />
          </div>
          <p class="text-sm">
            {{ favoritesCount }}
          </p>
        </div>
        <div
          class="flex items-center text-dark dark:text-light hover:text-darkblue dark:hover:text-darkblue"
        >
          <div
            class="mr-3 p-2 hover:bg-darkblue hover:bg-opacity-20 rounded-full"
          >
            <IconShare />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
