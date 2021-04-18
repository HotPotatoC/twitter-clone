<template>
  <!-- Create tweet dialog -->
  <Dialog :show="showCreateFormDialog" @close="showCreateFormDialog = false">
    <form @submit.prevent="addNewTweet" class="w-full">
      <textarea
        v-model="newTweet.content"
        placeholder="What's happening?"
        class="mt-3 w-full focus:outline-none dark:bg-black dark:text-light"
      />

      <div class="mt-4 text-right">
        <button
          type="submit"
          class="h-10 px-4 font-semibold focus:outline-none rounded-full"
          :class="
            contentIsEmpty
              ? ['bg-dark', 'text-light', 'cursor-default']
              : ['bg-blue', 'hover:bg-darkblue', 'text-lightest']
          "
          @click="showCreateFormDialog = false"
          :disabled="contentIsEmpty"
        >
          Tweet
        </button>
      </div>
    </form>
  </Dialog>

  <!-- Sidebar -->
  <div
    class="lg:w-1/5 border-r border-lighter dark:border-darker px-2 lg:px-8 py-2 flex flex-col justify-between"
  >
    <div>
      <button
        class="h-12 w-12 hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-20 text-3xl text-blue rounded-full"
      >
        <FontAwesome :icon="['fab', 'twitter']" />
      </button>
      <div>
        <router-link v-for="tab in tabs" :key="tab.id" :to="tab.to">
          <button
            @click="selectedTab = tab.id"
            class="flex items-center mr-auto mb-3 py-2 px-4 rounded-full hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-20 hover:text-blue focus:outline-none"
            :class="
              selectedTab === tab.id
                ? 'text-blue dark:text-blue'
                : 'dark:text-light'
            "
          >
            <FontAwesome
              :icon="[tab.iconPrefix, tab.icon]"
              class="text-left text-2xl mr-4"
            />
            <p class="text-left text-lg font-semibold hidden lg:block">
              {{ tab.label }}
            </p>
          </button>
        </router-link>
        <button
          class="text-lightest bg-blue rounded-full font-semibold focus:outline-none w-12 h-12 lg:w-full lg:h-auto p-3 hover:bg-darkblue"
          @click="showCreateFormDialog = true"
        >
          <p class="hidden lg:block">Tweet</p>
          <FontAwesome :icon="['fas', 'plus']" class="lg:hidden" />
        </button>
      </div>
    </div>
    <div class="lg:w-full relative">
      <button
        @click="showDropdown = !showDropdown"
        class="flex items-center w-full hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-10 rounded-full p-2 focus:outline-none"
      >
        <div class="lg:ml-4">
          <p class="text-sm font-bold leading-tight dark:text-lightest">
            {{ user.name }}
          </p>
          <p class="text-sm leading-tight dark:text-light">@{{ user.name }}</p>
        </div>
        <FontAwesome
          :icon="['fas', 'angle-down']"
          class="ml-auto text-lg dark:text-lightest"
        />
      </button>

      <div
        v-if="showDropdown"
        class="absolute overflow-hidden bottom-0 left-0 w-64 mb-16 rounded-2xl shadow-md border border-lighter dark:border-darker"
      >
        <button
          @click="showDropdown = false"
          class="flex items-center w-full hover:bg-lightest dark:hover:bg-darkest p-3 bg-white dark:bg-black focus:outline-none"
        >
          <div>
            <p class="text-sm font-bold leading-tight dark:text-lightest">
              {{ user.name }}
            </p>
            <p class="text-sm leading-tight dark:text-light">
              @{{ user.name }}
            </p>
          </div>
          <FontAwesome :icon="['fas', 'check']" class="ml-auto text-blue" />
        </button>
        <button
          @click="logout"
          class="w-full text-left hover:bg-lightest dark:bg-black dark:hover:bg-darkest border-t border-lighter dark:border-darker p-3 text-sm dark:text-lightest focus:outline-none"
        >
          Log Out @{{ user.name }}
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref, Ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ActionTypes as AuthActionTypes } from '../../modules/auth/store/actions'
import { ActionTypes as TweetActionTypes } from '../../modules/tweets/store/actions'
import { useStore } from '../../store'
import Dialog from '../common/Dialog.vue'

interface Tab {
  id: string
  icon: string
  iconPrefix: string
  label: string
  to: string
}

interface NewTweet {
  content: string | Ref<string>
}

export default defineComponent({
  name: 'ProfileSidebar',
  components: { Dialog },
  setup() {
    const tabs: Tab[] = [
      {
        id: 'home',
        icon: 'home',
        iconPrefix: 'fas',
        label: 'Home',
        to: '/home',
      },
      {
        id: 'explore',
        icon: 'hashtag',
        iconPrefix: 'fas',
        label: 'Explore',
        to: '/home',
      },
      {
        id: 'notifications',
        icon: 'bell',
        iconPrefix: 'fas',
        label: 'Notifications',
        to: '/home',
      },
      {
        id: 'messages',
        icon: 'envelope',
        iconPrefix: 'fas',
        label: 'Messages',
        to: '/home',
      },
      {
        id: 'bookmarks',
        icon: 'bookmark',
        iconPrefix: 'fas',
        label: 'Bookmarks',
        to: '/home',
      },
      {
        id: 'lists',
        icon: 'clipboard-list',
        iconPrefix: 'fas',
        label: 'Lists',
        to: '/home',
      },
      {
        id: 'profile',
        icon: 'user',
        iconPrefix: 'fas',
        label: 'Profile',
        to: '/home',
      },
      {
        id: 'more',
        icon: 'ellipsis-h',
        iconPrefix: 'fas',
        label: 'More',
        to: '/home',
      },
    ]

    const store = useStore()
    const router = useRouter()
    const selectedTab = ref<string>('home')
    const showDropdown = ref<boolean>(false)
    const showCreateFormDialog = ref<boolean>(false)
    const tweetContent = ref<string>('')
    const newTweet = reactive<NewTweet>({
      content: tweetContent,
    })

    const user = computed(() => store.getters['userData'])
    const contentIsEmpty = computed(() => tweetContent.value === '')

    async function addNewTweet() {
      try {
        await store.dispatch(TweetActionTypes.NEW_TWEET, newTweet.content)
        newTweet.content = ''
      } catch (error) {
        console.log(error)
      }
    }

    async function logout() {
      await store.dispatch(AuthActionTypes.LOGOUT_USER)

      router.push('/login')
      return
    }

    return {
      tabs,
      selectedTab,
      showDropdown,
      showCreateFormDialog,
      user,
      logout,
      tweetContent,
      newTweet,
      contentIsEmpty,
      addNewTweet,
    }
  },
})
</script>
