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

enum TweetTypeCode {
  TWEET,
  REPLY,
}

export default defineComponent({
  name: 'TweetConversationCard',
  components: { IconComment, IconEllipsisH, IconRetweet, IconHeart, IconShare },
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

    const replyFavoritesCount = ref(tweet.value.favoritesCount)
    const replyAlreadyLiked = ref(tweet.value.alreadyLiked)

    const repliedTweetFavoritesCount = ref(tweet.value.repliedTo.favoritesCount)
    const repliedTweetAlreadyLiked = ref(tweet.value.repliedTo.alreadyLiked)

    const parsedCreatedAt = computed(() => {
      dayjs.extend(relativeTime)
      return dayjs(tweet.value.createdAt).fromNow()
    })

    const parsedReplyContent = computed(() =>
      linkifyHTMLText(tweet.value.content)
    )

    const parsedRepliedTweetContent = computed(() =>
      linkifyHTMLText(tweet.value.repliedTo.content)
    )

    const replyAlreadyLikedClasses = computed(() => {
      return replyAlreadyLiked.value
        ? ['text-danger']
        : [
            'text-dark',
            'dark:text-light',
            'hover:text-danger',
            'dark:hover:text-danger',
          ]
    })

    const repliedTweetAlreadyLikedClasses = computed(() => {
      return repliedTweetAlreadyLiked.value
        ? ['text-danger']
        : [
            'text-dark',
            'dark:text-light',
            'hover:text-danger',
            'dark:hover:text-danger',
          ]
    })

    async function likeTweet(typeCode: TweetTypeCode, tweetId: string) {
      await store.dispatch(Action.TweetsActionTypes.FAVORITE_TWEET, tweetId)

      switch (typeCode) {
        case TweetTypeCode.TWEET:
          repliedTweetAlreadyLiked.value = !repliedTweetAlreadyLiked.value
          if (repliedTweetAlreadyLiked.value) {
            repliedTweetFavoritesCount.value++
          } else {
            repliedTweetFavoritesCount.value--
          }
          break
        case TweetTypeCode.REPLY:
          replyAlreadyLiked.value = !replyAlreadyLiked.value
          if (replyAlreadyLiked.value) {
            replyFavoritesCount.value++
          } else {
            replyFavoritesCount.value--
          }
          break
        default:
          console.error('Invalid operation type code')
          break
      }
    }

    function showOverlay(imageURL: string) {
      store.dispatch(Action.TweetsActionTypes.TOGGLE_TWEET_IMAGE_OVERLAY, {
        tweet: tweet.value,
        show: true,
        source: imageURL,
      })
    }

    return {
      router,
      replyAlreadyLiked,
      replyFavoritesCount,
      repliedTweetFavoritesCount,
      repliedTweetAlreadyLiked,
      parsedCreatedAt,
      parsedReplyContent,
      parsedRepliedTweetContent,
      replyAlreadyLikedClasses,
      repliedTweetAlreadyLikedClasses,
      likeTweet,
      TweetTypeCode,
      showOverlay,
    }
  },
})
</script>

<template>
  <div
    class="
      w-full
      p-4
      hover:bg-lighter
      dark:hover:bg-darker dark:hover:bg-opacity-30
      cursor-pointer
      transition-colors
      duration-75
    "
    @click="
      router.push(
        `/${tweet.repliedTo.authorHandle}/status/${tweet.repliedTo.id}`
      )
    "
  >
    <div class="flex">
      <div class="flex-none mr-4">
        <router-link :to="`/${tweet.repliedTo.authorHandle}`">
          <img
            v-lazy="tweet.repliedTo.authorPhotoURL"
            class="h-12 w-12 rounded-full flex-none"
          />
        </router-link>
        <div class="h-full w-1 ml-6 bg-dark"></div>
      </div>
      <div class="w-full">
        <div class="flex flex-wrap items-center w-full">
          <router-link
            :to="`/${tweet.repliedTo.authorHandle}`"
            class="flex flex-wrap items-center"
          >
            <p class="font-semibold dark:text-lightest hover:underline">
              {{ tweet.repliedTo.authorName }}
            </p>
            <p class="text-sm text-dark dark:text-light ml-2">
              @{{ tweet.repliedTo.authorHandle }} ·
            </p>
            <p class="text-sm text-dark dark:text-light ml-2">
              {{ parsedCreatedAt }}
            </p>
          </router-link>
          <div
            class="
              text-gray
              ml-auto
              p-2
              hover:bg-darkblue
              hover:text-blue
              hover:bg-opacity-20
              rounded-full
            "
          >
            <IconEllipsisH />
          </div>
        </div>
        <div
          class="py-2 break-words dark:text-lightest"
          v-html="parsedRepliedTweetContent"
        ></div>
        <div
          v-if="
            tweet.repliedTo.photoURLs !== null &&
            tweet.repliedTo.photoURLs.length > 0
          "
          class="relative overflow-hidden w-full h-96 rounded-lg"
        >
          <div class="box-border relative">
            <div
              v-if="tweet.repliedTo.photoURLs.length > 1"
              class="grid grid-cols-2 gap-1 h-full"
            >
              <div
                class="w-full"
                :class="tweet.repliedTo.photoURLs.length > 2 ? 'h-48' : 'h-96'"
              >
                <img
                  v-lazy="tweet.repliedTo.photoURLs[0]"
                  @click.stop="showOverlay(tweet.repliedTo.photoURLs[0])"
                  class="object-cover w-full h-full"
                />
                <img
                  v-if="tweet.repliedTo.photoURLs.length > 2"
                  v-lazy="tweet.repliedTo.photoURLs[1]"
                  @click.stop="showOverlay(tweet.repliedTo.photoURLs[1])"
                  class="object-cover w-full h-full"
                />
              </div>
              <div
                class="w-full"
                :class="tweet.repliedTo.photoURLs.length > 2 ? 'h-48' : 'h-96'"
              >
                <img
                  v-if="tweet.repliedTo.photoURLs.length > 2"
                  v-lazy="tweet.repliedTo.photoURLs[2]"
                  @click.stop="showOverlay(tweet.repliedTo.photoURLs[2])"
                  class="object-cover w-full"
                  :class="
                    tweet.repliedTo.photoURLs.length === 4 ? 'h-full' : 'h-96'
                  "
                />
                <img
                  v-else
                  v-lazy="tweet.repliedTo.photoURLs[1]"
                  @click.stop="showOverlay(tweet.repliedTo.photoURLs[1])"
                  class="object-cover w-full h-full"
                />
                <img
                  v-if="tweet.repliedTo.photoURLs.length === 4"
                  v-lazy="tweet.repliedTo.photoURLs[3]"
                  @click.stop="showOverlay(tweet.repliedTo.photoURLs[3])"
                  class="object-cover w-full h-full"
                />
              </div>
            </div>
            <div v-else class="w-full">
              <img
                v-lazy="tweet.repliedTo.photoURLs[0]"
                @click.stop="showOverlay(tweet.repliedTo.photoURLs[0])"
                class="object-cover w-full h-96"
              />
            </div>
          </div>
        </div>
        <div class="flex items-center justify-between w-full mt-2">
          <div
            class="
              flex
              items-center
              group
              text-dark
              dark:text-light
              hover:text-blue
              dark:hover:text-blue
            "
          >
            <div
              class="
                mr-3
                p-2
                group-hover:bg-darkblue group-hover:bg-opacity-20
                rounded-full
              "
            >
              <IconComment />
            </div>
            <p class="text-sm">
              {{ tweet.repliedTo.repliesCount }}
            </p>
          </div>
          <div
            class="
              flex
              items-center
              group
              text-dark
              dark:text-light
              hover:text-success
              dark:hover:text-success
            "
          >
            <div
              class="
                mr-3
                p-2
                group-hover:bg-success group-hover:bg-opacity-20
                rounded-full
              "
            >
              <IconRetweet />
            </div>
            <p class="text-sm">
              {{ tweet.repliedTo.repliesCount }}
            </p>
          </div>
          <div
            class="flex items-center group"
            :class="repliedTweetAlreadyLikedClasses"
          >
            <div
              class="
                mr-3
                p-2
                group-hover:bg-danger group-hover:bg-opacity-20
                rounded-full
              "
              @click.stop="
                likeTweet(TweetTypeCode.TWEET, tweet.repliedTo.id.toString())
              "
            >
              <IconHeart
                :class="repliedTweetAlreadyLiked ? 'fill-current' : null"
              />
            </div>
            <p class="text-sm">
              {{ repliedTweetFavoritesCount }}
            </p>
          </div>
          <div
            class="
              flex
              items-center
              text-dark
              dark:text-light
              hover:text-darkblue
              dark:hover:text-darkblue
            "
          >
            <div
              class="
                mr-3
                p-2
                hover:bg-darkblue hover:bg-opacity-20
                rounded-full
              "
            >
              <IconShare />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="
      w-full
      p-4
      border-b border-lighter
      dark:border-dark
      hover:bg-lighter
      dark:hover:bg-darker dark:hover:bg-opacity-30
      cursor-pointer
      transition-colors
      duration-75
    "
    @click="router.push(`/${tweet.authorHandle}/status/${tweet.id}`)"
  >
    <div class="flex">
      <router-link :to="`/${tweet.authorHandle}`" class="flex-none mr-4">
        <img
          v-lazy="tweet.authorPhotoURL"
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
            class="
              text-gray
              ml-auto
              p-2
              hover:bg-darkblue
              hover:text-blue
              dark:hover:text-blue
              hover:bg-opacity-20
              rounded-full
            "
          >
            <IconEllipsisH />
          </div>
        </div>
        <div class="pb-2 break-words">
          <span class="text-gray">Replying to</span>
          <router-link :to="`/${tweet.repliedTo.authorHandle}`">
            <span class="text-blue hover:underline">
              @{{ tweet.repliedTo.authorHandle }}
            </span>
          </router-link>
        </div>
        <div
          class="py-2 break-words dark:text-lightest"
          v-html="parsedReplyContent"
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
                  v-lazy="tweet.photoURLs[0]"
                  @click.stop="showOverlay(tweet.photoURLs[0])"
                  class="object-cover w-full h-full"
                />
                <img
                  v-if="tweet.photoURLs.length > 2"
                  v-lazy="tweet.photoURLs[1]"
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
                  v-lazy="tweet.photoURLs[2]"
                  @click.stop="showOverlay(tweet.photoURLs[2])"
                  class="object-cover w-full"
                  :class="tweet.photoURLs.length === 4 ? 'h-full' : 'h-96'"
                />
                <img
                  v-else
                  v-lazy="tweet.photoURLs[1]"
                  @click.stop="showOverlay(tweet.photoURLs[1])"
                  class="object-cover w-full h-full"
                />
                <img
                  v-if="tweet.photoURLs.length === 4"
                  v-lazy="tweet.photoURLs[3]"
                  @click.stop="showOverlay(tweet.photoURLs[3])"
                  class="object-cover w-full h-full"
                />
              </div>
            </div>
          </div>
        </div>
        <div class="flex items-center justify-between w-full mt-2">
          <div
            class="
              flex
              items-center
              group
              text-dark
              dark:text-light
              hover:text-blue
              dark:hover:text-blue
            "
          >
            <div
              class="
                mr-3
                p-2
                group-hover:bg-darkblue group-hover:bg-opacity-20
                rounded-full
              "
            >
              <IconComment />
            </div>
            <p class="text-sm">
              {{ tweet.repliesCount }}
            </p>
          </div>
          <div
            class="
              flex
              items-center
              group
              text-dark
              dark:text-light
              hover:text-success
              dark:hover:text-success
            "
          >
            <div
              class="
                mr-3
                p-2
                group-hover:bg-success group-hover:bg-opacity-20
                rounded-full
              "
            >
              <IconRetweet />
            </div>
            <p class="text-sm">
              {{ tweet.repliesCount }}
            </p>
          </div>
          <div
            class="flex items-center group"
            :class="replyAlreadyLikedClasses"
          >
            <div
              class="
                mr-3
                p-2
                group-hover:bg-danger group-hover:bg-opacity-20
                rounded-full
              "
              @click.stop="likeTweet(TweetTypeCode.REPLY, tweet.id.toString())"
            >
              <IconHeart :class="replyAlreadyLiked ? 'fill-current' : null" />
            </div>
            <p class="text-sm">
              {{ replyFavoritesCount }}
            </p>
          </div>
          <div
            class="
              flex
              items-center
              text-dark
              dark:text-light
              hover:text-darkblue
              dark:hover:text-darkblue
            "
          >
            <div
              class="
                mr-3
                p-2
                hover:bg-darkblue hover:bg-opacity-20
                rounded-full
              "
            >
              <IconShare />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
