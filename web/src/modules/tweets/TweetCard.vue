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

export default defineComponent({
  name: 'TweetCard',
  components: { IconComment, IconEllipsisH, IconRetweet, IconHeart, IconShare },
  props: {
    tweet: {
      type: Object as () => Tweet,
      required: true,
    },
  },
  setup(props) {
    const store = useStore()
    const { tweet } = toRefs(props)
    const favoritesCount = ref(tweet.value.favoritesCount)
    const alreadyLiked = ref(tweet.value.alreadyLiked)

    const parsedCreatedAt = computed(() => {
      dayjs.extend(relativeTime)
      return dayjs(tweet.value.createdAt).fromNow()
    })

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

    return { alreadyLiked, favoritesCount, parsedCreatedAt, likeTweet }
  },
})
</script>

<template>
  <div
    class="w-full p-4 border-b border-lighter dark:border-dark hover:bg-lighter dark:hover:bg-darker dark:hover:bg-opacity-30 flex cursor-pointer transition-colors duration-75"
  >
    <router-link :to="`/${tweet.authorHandle}`" class="flex-none mr-4">
      <img
        :src="tweet.authorPhotoURL"
        class="h-12 w-12 rounded-full flex-none"
      />
    </router-link>
    <div class="w-full">
      <router-link :to="`/${tweet.authorHandle}/status/${tweet.id}`">
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
        <p class="py-2 break-words dark:text-lightest">
          {{ tweet.content }}
        </p>
      </router-link>
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
            @click="likeTweet"
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
