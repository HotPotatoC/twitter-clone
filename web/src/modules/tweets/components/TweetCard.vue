<script lang="ts">
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import { computed, defineComponent } from 'vue'
import { Tweet } from '../store/state'
import IconEllipsisH from '../../../components/icons/IconEllipsisH.vue'
import IconComment from '../../../components/icons/IconComment.vue'
import IconRetweet from '../../../components/icons/IconRetweet.vue'
import IconHeart from '../../../components/icons/IconHeart.vue'
import IconShare from '../../../components/icons/IconShare.vue'

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
    const parsedCreatedAt = computed(() => {
      dayjs.extend(relativeTime)
      return dayjs(props.tweet.createdAt).fromNow()
    })

    return { parsedCreatedAt }
  },
})
</script>

<template>
  <div class="flex flex-wrap items-center w-full">
    <p class="font-semibold dark:text-lightest">{{ tweet.name }}</p>
    <p class="text-sm text-dark dark:text-light ml-2">@{{ tweet.handle }} ·</p>
    <p class="text-sm text-dark dark:text-light ml-2">
      {{ parsedCreatedAt }}
    </p>
    <div
      class="text-gray ml-auto p-2 hover:bg-darkblue hover:text-blue hover:bg-opacity-20 rounded-full"
    >
      <IconEllipsisH />
    </div>
  </div>
  <p class="py-2 break-words dark:text-lightest">
    {{ tweet.content }}
  </p>
  <div class="flex items-center justify-between w-full mt-2">
    <div class="flex items-center">
      <div
        class="mr-3 p-2 text-dark dark:text-light hover:bg-darkblue hover:text-blue hover:bg-opacity-20 rounded-full"
      >
        <IconComment />
      </div>
      <p class="text-sm text-dark dark:text-light">
        {{ tweet.repliesCount }}
      </p>
    </div>
    <div class="flex items-center">
      <div
        class="mr-3 p-2 text-dark dark:text-light hover:bg-success hover:text-success hover:bg-opacity-20 rounded-full"
      >
        <IconRetweet />
      </div>
      <p class="text-sm text-dark dark:text-light">
        {{ tweet.repliesCount }}
      </p>
    </div>
    <div class="flex items-center">
      <div
        class="mr-3 p-2 hover:bg-danger hover:bg-opacity-20 rounded-full"
        :class="
          tweet.alreadyLiked
            ? ['text-danger']
            : ['text-dark', 'dark:text-light', 'hover:text-danger']
        "
      >
        <IconHeart :class="tweet.alreadyLiked ? 'fill-current' : null" />
      </div>
      <p class="text-sm text-dark dark:text-light">
        {{ tweet.favoritesCount }}
      </p>
    </div>
    <div class="flex items-center">
      <div
        class="mr-3 p-2 text-dark dark:text-light hover:bg-darkblue hover:text-darkblue hover:bg-opacity-20 rounded-full"
      >
        <IconShare />
      </div>
    </div>
  </div>
</template>
