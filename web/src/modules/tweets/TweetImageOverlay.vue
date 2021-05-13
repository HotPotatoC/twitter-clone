<script lang="ts">
import { defineComponent, toRefs, computed } from 'vue'
import dayjs from 'dayjs'
import {
  TransitionRoot,
  TransitionChild,
  Dialog,
  DialogOverlay,
  DialogTitle,
  DialogDescription,
} from '@headlessui/vue'
import IconX from '../../icons/IconX.vue'
import IconEllipsisH from '../../icons/IconEllipsisH.vue'
import IconComment from '../../icons/IconComment.vue'
import IconRetweet from '../../icons/IconRetweet.vue'
import IconShare from '../../icons/IconShare.vue'
import IconHeart from '../../icons/IconHeart.vue'
import { Tweet } from './types'
import { useStore } from '../../store'
import { linkifyHTMLText } from '../../utils/linkify'
import { Action } from '../storeActionTypes'

export default defineComponent({
  name: 'TweetImageOverlay',
  components: {
    TransitionRoot,
    TransitionChild,
    Dialog,
    DialogOverlay,
    DialogTitle,
    DialogDescription,
    IconX,
    IconEllipsisH,
    IconComment,
    IconRetweet,
    IconHeart,
    IconShare,
  },
  props: {
    tweet: {
      type: Object as () => Tweet,
      required: true,
    },
    show: {
      type: Boolean,
      default: false,
    },
    image: {
      type: String,
      required: true,
    },
  },
  emits: ['close'],
  setup(props, { emit }) {
    const { tweet, show, image } = toRefs(props)

    const store = useStore()

    const parsedCreatedAt = computed(() =>
      dayjs(tweet.value.createdAt).format('h:mm A · MMM D, YYYY')
    )

    const parsedContent = computed(() => linkifyHTMLText(tweet.value.content))

    async function likeTweet() {
      await store.dispatch(
        Action.TweetsActionTypes.FAVORITE_TWEET,
        tweet.value.id.toString()
      )

      tweet.value.alreadyLiked = !tweet.value.alreadyLiked
      if (tweet.value.alreadyLiked) {
        tweet.value.favoritesCount++
      } else {
        tweet.value.favoritesCount--
      }
    }

    function onClose() {
      emit('close')
    }

    return {
      tweet,
      show,
      onClose,
      image,
      parsedCreatedAt,
      parsedContent,
      likeTweet,
    }
  },
})
</script>

<template>
  <TransitionRoot appear :show="show" as="template">
    <Dialog as="div" static :open="show" @close="onClose">
      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="min-h-screen">
          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0"
            enter-to="opacity-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100"
            leave-to="opacity-0"
          >
            <DialogOverlay class="fixed inset-0 bg-blue bg-opacity-20" />
          </TransitionChild>

          <span class="inline-block h-screen align-middle" aria-hidden="true">
            &#8203;
          </span>

          <TransitionChild
            as="template"
            enter="duration-100 ease-out"
            enter-from="opacity-0"
            enter-to="opacity-100"
            leave="duration-100 ease-in"
            leave-from="opacity-100"
            leave-to="opacity-0"
          >
            <div
              class="
                inline-block
                w-full
                h-full
                overflow-hidden
                align-middle
                transition-all
                transform
              "
            >
              <button
                @click="onClose()"
                class="
                  absolute
                  top-0
                  left-0
                  m-4
                  p-1
                  border-none
                  rounded-full
                  focus:outline-none
                  hover:bg-white hover:bg-opacity-20
                "
              >
                <IconX :size="26" class="dark:text-white" />
              </button>
              <div class="flex items-center h-full">
                <div class="w-full flex justify-center">
                  <img
                    v-lazy="image"
                    class="max-w-lg md:max-w-xl xl:max-w-5xl"
                  />
                </div>
                <div
                  class="
                    w-1/4
                    h-screen
                    overflow-y-scroll
                    bg-black
                    border-l border-lighter
                    dark:border-darker
                  "
                >
                  <div class="p-4 border-b border-lighter dark:border-dark">
                    <div class="w-full">
                      <div class="flex items-center w-full">
                        <img
                          v-lazy="tweet.authorPhotoURL"
                          class="mr-5 h-12 w-12 rounded-full"
                        />
                        <router-link :to="`/${tweet.authorHandle}`">
                          <p
                            class="
                              font-semibold
                              dark:text-lightest
                              hover:underline
                            "
                          >
                            {{ tweet.authorName }}
                          </p>
                          <p class="text-sm text-dark dark:text-light">
                            @{{ tweet.authorHandle }}
                          </p>
                        </router-link>
                        <div
                          class="
                            cursor-pointer
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
                        class="text-xl py-4 break-words dark:text-lightest"
                        v-html="parsedContent"
                      ></div>
                      <p class="text-dark dark:text-gray">
                        {{ parsedCreatedAt }}
                      </p>
                      <div
                        class="
                          flex
                          items-center
                          justify-start
                          space-x-12
                          w-full
                          border-t border-b
                          my-4
                          py-4
                          border-lighter
                          dark:border-dark
                        "
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
                        class="
                          flex
                          items-center
                          justify-around
                          w-full
                          text-xl text-dark
                          dark:text-light
                        "
                      >
                        <div
                          class="
                            flex
                            justify-center
                            hover:bg-darkblue
                            hover:text-blue
                            hover:bg-opacity-20
                            rounded-full
                            p-3
                            cursor-pointer
                            transition
                            duration-75
                          "
                          @click="showCreateReplyDialog = true"
                        >
                          <IconComment :size="20" />
                        </div>
                        <div
                          class="
                            flex
                            justify-center
                            hover:bg-success
                            hover:text-success
                            hover:bg-opacity-20
                            rounded-full
                            p-3
                            cursor-pointer
                          "
                        >
                          <IconRetweet :size="20" />
                        </div>
                        <div
                          class="
                            flex
                            justify-center
                            hover:bg-danger
                            hover:text-danger
                            hover:bg-opacity-20
                            rounded-full
                            p-3
                            cursor-pointer
                          "
                          :class="
                            tweet.alreadyLiked
                              ? ['text-danger']
                              : [
                                  'text-dark',
                                  'dark:text-light',
                                  'hover:text-danger',
                                ]
                          "
                          @click="likeTweet"
                        >
                          <IconHeart
                            :size="20"
                            :class="tweet.alreadyLiked ? 'fill-current' : null"
                          />
                        </div>
                        <div
                          class="
                            flex
                            justify-center
                            hover:bg-darkblue
                            hover:text-darkblue
                            hover:bg-opacity-20
                            rounded-full
                            p-3
                            cursor-pointer
                          "
                        >
                          <IconShare :size="20" />
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
