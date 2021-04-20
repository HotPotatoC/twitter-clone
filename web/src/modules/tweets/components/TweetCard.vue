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
    <IconEllipsisH class="text-gray ml-auto" />
  </div>
  <p class="py-2 break-words dark:text-lightest">
    {{ tweet.content }}
  </p>
  <div class="flex items-center justify-between w-full">
    <div class="flex items-center text-sm text-dark dark:text-light">
      <IconComment class="mr-3" />
      <p>{{ tweet.repliesCount }}</p>
    </div>
    <div class="flex items-center text-sm text-dark dark:text-light">
      <IconRetweet class="mr-3" />
      <p>{{ tweet.repliesCount }}</p>
    </div>
    <div class="flex items-center text-sm text-dark dark:text-light">
      <IconHeart class="mr-3" />
      <p>{{ tweet.favoritesCount }}</p>
    </div>
    <div class="flex items-center text-sm text-dark dark:text-light">
      <IconShare class="mr-3" />
    </div>
  </div>
</template>
