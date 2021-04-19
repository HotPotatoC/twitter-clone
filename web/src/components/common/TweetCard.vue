<template>
  <div class="flex items-center w-full">
    <p class="font-semibold dark:text-lightest">{{ tweet.name }}</p>
    <p class="text-sm text-dark dark:text-light ml-2">@{{ tweet.name }} ·</p>
    <p class="text-sm text-dark dark:text-light ml-2">
      {{ parsedCreatedAt }}
    </p>
    <FontAwesome :icon="['fas', 'ellipsis-h']" class="text-gray ml-auto" />
  </div>
  <p class="py-2 break-words dark:text-lightest">
    {{ tweet.content }}
  </p>
  <div class="flex items-center justify-between w-full">
    <div class="flex items-center text-sm text-dark dark:text-light">
      <FontAwesome :icon="['fas', 'comment']" class="mr-3" />
      <p>{{ tweet.repliesCount }}</p>
    </div>
    <div class="flex items-center text-sm text-dark dark:text-light">
      <FontAwesome :icon="['fas', 'retweet']" class="mr-3" />
      <p>{{ tweet.repliesCount }}</p>
    </div>
    <div class="flex items-center text-sm text-dark dark:text-light">
      <FontAwesome :icon="['fas', 'heart']" class="mr-3" />
      <p>{{ tweet.favoritesCount }}</p>
    </div>
    <div class="flex items-center text-sm text-dark dark:text-light">
      <FontAwesome :icon="['fas', 'share-square']" class="mr-3" />
    </div>
  </div>
</template>

<script lang="ts">
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import { computed, defineComponent } from 'vue'
import { Tweet } from '../../modules/tweets/store/state'

export default defineComponent({
  name: 'TweetCard',
  props: {
    tweet: {
      type: Object as () => Tweet,
      required: true,
    },
  },
  setup(props) {
    const parsedCreatedAt = computed(() => {
      dayjs.extend(relativeTime)
      return dayjs(props.tweet.createdAt).fromNow()
    })

    return { parsedCreatedAt }
  },
})
</script>
