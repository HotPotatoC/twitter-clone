<template>
  <main
    class="w-full h-full overflow-y-scroll border-r border-lighter dark:border-darker md:border-r-0"
    ref="tweetsRef"
    @scroll="handleScroll"
  >
    <div
      class="px-5 py-3 border-b border-lighter dark:border-dark flex items-center justify-start space-x-6"
    >
      <Return />
      <h1 class="text-xl font-bold dark:text-lightest">{{ profile.name }}</h1>
    </div>
    <div
      class="px-6 h-48 border-b border-lighter dark:border-dark bg-blue relative"
    >
      <div
        class="absolute mt-32 rounded-full w-32 h-32 bg-dark border-4 border-lightest dark:border-black"
      ></div>
    </div>
    <div class="mt-5 px-6">
      <button
        v-show="isCurrentUser"
        class="float-right text-blue font-bold py-2 px-5 rounded-full border-2 border-blue hover:bg-blue hover:bg-opacity-25 focus:outline-none transition-colors duration-75"
      >
        Edit Profile
      </button>
      <button
        v-show="!isCurrentUser"
        class="float-right text-blue font-bold py-2 px-6 rounded-full border-2 border-blue hover:bg-blue hover:bg-opacity-25 focus:outline-none transition-colors duration-75"
      >
        Follow
      </button>
      <button
        v-show="!isCurrentUser"
        class="float-right text-blue font-bold py-2 px-3 mr-4 rounded-full border-2 border-blue hover:bg-blue hover:bg-opacity-25 focus:outline-none transition-colors duration-75"
      >
        <FontAwesome :icon="['fas', 'ellipsis-h']" />
      </button>
      <div class="w-full flex flex-col space-y-4">
        <div class="w-full block mt-4">
          <h1 class="text-2xl font-bold dark:text-lightest">
            {{ profile.name }}
          </h1>
          <h2 class="dark:text-gray">@{{ profile.name }}</h2>
        </div>
        <p class="w-full dark:text-lightest">
          {{ profile.bio }}
        </p>
        <div
          class="w-full flex flex-wrap flex-col md:flex-row md:space-x-4 text-sm"
        >
          <div
            v-show="profile.location !== ''"
            class="flex items-center space-x-1 dark:text-gray"
          >
            <FontAwesome :icon="['fas', 'map-marker-alt']" />
            <p>{{ profile.location }}</p>
          </div>
          <div
            v-show="profile.website !== ''"
            class="flex items-center space-x-1 dark:text-gray"
          >
            <FontAwesome :icon="['fas', 'link']" />
            <a
              :href="profile.website"
              target="_blank"
              rel="noreferrer noopener"
              class="text-blue hover:underline"
            >
              {{ profile.website }}
            </a>
          </div>
          <div
            v-show="validBirthDate"
            class="flex items-center space-x-1 dark:text-gray"
          >
            <FontAwesome :icon="['fas', 'birthday-cake']" />
            <p>Born {{ parsedBirthDate }}</p>
          </div>
          <div class="flex items-center space-x-1 dark:text-gray">
            <FontAwesome :icon="['fas', 'calendar-alt']" />
            <p>Joined {{ parsedJoinedAt }}</p>
          </div>
        </div>
        <div
          class="w-full flex flex-wrap flex-col md:flex-row md:space-x-4 text-sm"
        >
          <div class="flex items-center space-x-2">
            <p class="font-bold dark:text-lightest">
              {{ profile.followingsCount }}
            </p>
            <p class="dark:text-gray">Following</p>
          </div>
          <div class="flex items-center space-x-2">
            <p class="font-bold dark:text-lightest">
              {{ profile.followersCount }}
            </p>
            <p class="dark:text-gray">Followers</p>
          </div>
        </div>
      </div>
    </div>
    <div
      class="w-full flex justify-between items-center mt-4 border-b border-lighter dark:border-dark"
    >
      <div
        class="w-full py-4 text-center hover:bg-blue hover:bg-opacity-10 hover:text-blue cursor-pointer transition-colors duration-75"
        :class="selectedTab === 1 ? tabClasses : selectedTabClasses"
        @click="selectedTab = 1"
      >
        <h1 class="font-bold">Tweets</h1>
      </div>
      <div
        class="w-full py-4 text-center hover:bg-blue hover:bg-opacity-10 hover:text-blue cursor-pointer transition-colors duration-75"
        :class="selectedTab === 2 ? tabClasses : selectedTabClasses"
        @click="selectedTab = 2"
      >
        <h1 class="font-bold">Tweets & Replies</h1>
      </div>
      <div
        class="w-full py-4 text-center hover:bg-blue hover:bg-opacity-10 hover:text-blue cursor-pointer transition-colors duration-75"
        :class="selectedTab === 3 ? tabClasses : selectedTabClasses"
        @click="selectedTab = 3"
      >
        <h1 class="font-bold">Media</h1>
      </div>
      <div
        class="w-full py-4 text-center hover:bg-blue hover:bg-opacity-10 hover:text-blue cursor-pointer transition-colors duration-75"
        :class="selectedTab === 4 ? tabClasses : selectedTabClasses"
        @click="selectedTab = 4"
      >
        <h1 class="font-bold">Likes</h1>
      </div>
    </div>

    <div v-show="initialLoadDone" class="flex flex-col">
      <div
        v-for="tweet in tweets"
        :key="tweet.id"
        class="w-full p-4 border-b dark:border-dark hover:bg-lighter dark:hover:bg-darker flex cursor-pointer transition-colors duration-75"
      >
        <div class="w-full">
          <router-link :to="`/${tweet.name}/status/${tweet.id}`">
            <TweetCard :tweet="tweet" />
          </router-link>
        </div>
      </div>

      <div
        v-show="tweets.length > 0 && loadNextBatch"
        class="w-full p-4 border-b dark:border-dark hover:bg-lighter dark:hover:bg-darker flex cursor-pointer"
      >
        <div class="w-full text-center">
          <LoadingSpinner />
        </div>
      </div>
    </div>
  </main>
</template>

<script lang="ts">
import {
  defineComponent,
  onBeforeMount,
  onMounted,
  ref,
  computed,
  watch,
} from 'vue'
import { useRoute } from 'vue-router'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import { useStore } from '../../store'
import { ActionTypes } from './store/actions'
import Return from '../../components/common/Return.vue'
import TweetCard from '../tweets/components/TweetCard.vue'
import { Tweet } from '../tweets/store/state'
import LoadingSpinner from '../../components/common/LoadingSpinner.vue'

export default defineComponent({
  name: 'Profile',
  components: { Return, TweetCard, LoadingSpinner },
  setup() {
    const store = useStore()
    const route = useRoute()

    const selectedTab = ref<number>(1)
    const tweetsRef = ref<Element>(null)
    const initialLoadDone = ref<boolean>(false)
    const loadNextBatch = ref<boolean>(false)
    const tweets = ref<Tweet[]>([])

    const tabClasses = ['text-blue', 'border-b-2', 'border-blue']

    const selectedTabClasses = [
      'dark:text-gray',
      'border-b-2',
      'border-dark',
      'border-opacity-0',
    ]

    const profile = computed(() => store.getters['profileInfo'])

    const isCurrentUser = computed(
      () => store.getters['userData'].id === store.getters['profileInfo'].id
    )

    // Check if profile birthdate is set or not because it is set to "" by default.
    const validBirthDate = computed(() => {
      dayjs.extend(relativeTime)
      return dayjs(store.getters['profileInfo'].birthDate).isValid()
    })

    const parsedBirthDate = computed(() => {
      dayjs.extend(relativeTime)
      return dayjs(store.getters['profileInfo'].birthDate).format(
        'MMMM D, YYYY'
      )
    })

    const parsedJoinedAt = computed(() => {
      dayjs.extend(relativeTime)
      return dayjs(store.getters['profileInfo'].joinedAt).format('MMMM YYYY')
    })

    onBeforeMount(async () => {
      await store.dispatch(ActionTypes.GET_PROFILE_DETAILS, route.params.name)
      await loadTweets()
      initialLoadDone.value = true
    })

    onMounted(() => {
      watch(
        () => route.params.name,
        async (name) => {
          initialLoadDone.value = false
          await store.dispatch(ActionTypes.GET_PROFILE_DETAILS, name)
          await loadTweets()
          initialLoadDone.value = true
        },
        { flush: 'post' }
      )
    })

    async function loadTweets() {
      if (initialLoadDone.value && store.getters['profileTweets'].length > 0) {
        const lastItem = store.getters['lastProfileTweet']
        await store.dispatch(ActionTypes.LOAD_MORE_PROFILE_TWEETS, {
          username: route.params.name,
          cursor: lastItem.createdAt,
        })
      } else {
        await store.dispatch(ActionTypes.GET_PROFILE_TWEETS, route.params.name)
      }

      tweets.value = store.getters['profileTweets']
    }

    async function handleScroll(e: Event) {
      const element = tweetsRef.value
      if (
        !loadNextBatch.value &&
        element.scrollTop + element.clientHeight + 1 >= element.scrollHeight
      ) {
        loadNextBatch.value = true
        await loadTweets()
        loadNextBatch.value = false
      }
    }

    return {
      selectedTab,
      tabClasses,
      selectedTabClasses,
      tweetsRef,
      initialLoadDone,
      loadNextBatch,
      profile,
      tweets,
      isCurrentUser,
      validBirthDate,
      parsedBirthDate,
      parsedJoinedAt,
      handleScroll,
    }
  },
})
</script>
