<script lang="ts">
import { computed, defineComponent, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from '../../store'
import { ActionTypes as AuthActionTypes } from '../../modules/auth/store/actions'
import { ActionTypes as TweetActionTypes } from '../../modules/tweets/store/actions'
import Dialog from '../common/Dialog.vue'
import TweetCreateTweetDialog from '../../modules/tweets/components/TweetCreateTweetDialog.vue'

interface Tab {
  id: string
  icon: string
  iconPrefix: string
  label: string
  to: string
}

export default defineComponent({
  name: 'ProfileSidebar',
  components: { Dialog, TweetCreateTweetDialog },
  setup() {
    const store = useStore()
    const router = useRouter()
    const ready = ref<boolean>(false)
    const selectedTab = ref<string>('home')
    const showDropdown = ref<boolean>(false)
    const showCreateFormDialog = ref<boolean>(false)
    const tabs = ref<Tab[]>([
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
        to: '',
      },
      {
        id: 'more',
        icon: 'ellipsis-h',
        iconPrefix: 'fas',
        label: 'More',
        to: '/home',
      },
    ])

    const user = computed(() => store.getters['userData'])

    onMounted(() => {
      const index = tabs.value.findIndex((tab) => tab.id === 'profile')
      // Use watch to wait for the user data mutation to finish
      // then show the navigation tabs
      watch(
        () => store.getters['userData'].name,
        (value) => {
          // Update the link for profile tab to current user data name
          tabs.value[index].to = `/${value}`
          ready.value = true
        }
      )

      // Just in case
      ready.value = true
    })

    async function createTweet(content: string) {
      try {
        await store.dispatch(TweetActionTypes.NEW_TWEET, content)
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
      ready,
      selectedTab,
      showDropdown,
      showCreateFormDialog,
      user,
      logout,
      createTweet,
      tabs,
    }
  },
})
</script>

<template>
  <TweetCreateTweetDialog
    :show="showCreateFormDialog"
    @close="showCreateFormDialog = false"
    @dispatch="createTweet"
  />

  <div
    class="lg:w-1/5 border-r border-lighter dark:border-dark px-2 lg:px-8 py-2 flex flex-col justify-between"
  >
    <div v-show="ready">
      <button
        class="h-12 w-12 hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-20 text-3xl text-blue rounded-full transition-colors duration-75"
      >
        <FontAwesome :icon="['fab', 'twitter']" />
      </button>
      <div>
        <router-link v-for="tab in tabs" :key="tab.id" :to="tab.to">
          <button
            @click="selectedTab = tab.id"
            class="flex items-center mr-auto mb-3 py-2 px-4 rounded-full hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-20 hover:text-blue focus:outline-none transition-colors duration-75"
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
          class="text-lightest bg-blue rounded-full font-semibold focus:outline-none w-12 h-12 lg:w-full lg:h-auto p-3 hover:bg-darkblue transition-colors duration-75"
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
        class="flex items-center w-full hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-10 rounded-full p-2 focus:outline-none transition-colors duration-75"
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
        class="absolute overflow-hidden bottom-0 left-0 w-64 mb-16 rounded-2xl shadow-md border border-lighter dark:border-dark"
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
          class="w-full text-left hover:bg-lightest dark:bg-black dark:hover:bg-darkest border-t border-lighter dark:border-dark p-3 text-sm dark:text-lightest focus:outline-none"
        >
          Log Out @{{ user.name }}
        </button>
      </div>
    </div>
  </div>
</template>
