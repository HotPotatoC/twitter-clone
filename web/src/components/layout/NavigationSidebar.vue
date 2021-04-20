<script lang="ts">
import { computed, defineComponent, onMounted, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from '../../store'
import { ActionTypes as AuthActionTypes } from '../../modules/auth/store/actions'
import { ActionTypes as TweetActionTypes } from '../../modules/tweets/store/actions'
import Dialog from '../common/Dialog.vue'
import TweetCreateTweetDialog from '../../modules/tweets/components/TweetCreateTweetDialog.vue'
import IconTwitter from '../icons/IconTwitter.vue'
import IconPlus from '../icons/IconPlus.vue'
import IconArrowDown from '../icons/IconArrowDown.vue'
import IconChevronDown from '../icons/IconChevronDown.vue'
import IconCheck from '../icons/IconCheck.vue'

interface Tab {
  id: string
  icon: string
  label: string
  to: string
}

export default defineComponent({
  name: 'ProfileSidebar',
  components: {
    Dialog,
    TweetCreateTweetDialog,
    IconTwitter,
    IconPlus,
    IconChevronDown,
    IconCheck,
  },
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
        icon: `
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="28"
    height="28"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-home"
  >
    <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
    <polyline points="9 22 9 12 15 12 15 22"></polyline>
  </svg>`,
        label: 'Home',
        to: '/home',
      },
      {
        id: 'explore',
        icon: `
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="28"
    height="28"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-hash"
  >
    <line x1="4" y1="9" x2="20" y2="9"></line>
    <line x1="4" y1="15" x2="20" y2="15"></line>
    <line x1="10" y1="3" x2="8" y2="21"></line>
    <line x1="16" y1="3" x2="14" y2="21"></line>
  </svg>`,
        label: 'Explore',
        to: '/home',
      },
      {
        id: 'notifications',
        icon: `
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="28"
    height="28"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-bell"
  >
    <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
    <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
  </svg>`,
        label: 'Notifications',
        to: '/home',
      },
      {
        id: 'messages',
        icon: `
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="28"
    height="28"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-mail"
  >
    <path
      d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"
    ></path>
    <polyline points="22,6 12,13 2,6"></polyline>
  </svg>`,
        label: 'Messages',
        to: '/home',
      },
      {
        id: 'bookmarks',
        icon: `
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="28"
    height="28"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-bookmark"
  >
    <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
  </svg>`,
        label: 'Bookmarks',
        to: '/home',
      },
      {
        id: 'lists',
        icon: `
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="28"
    height="28"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-list"
  >
    <line x1="8" y1="6" x2="21" y2="6"></line>
    <line x1="8" y1="12" x2="21" y2="12"></line>
    <line x1="8" y1="18" x2="21" y2="18"></line>
    <line x1="3" y1="6" x2="3.01" y2="6"></line>
    <line x1="3" y1="12" x2="3.01" y2="12"></line>
    <line x1="3" y1="18" x2="3.01" y2="18"></line>
  </svg>`,
        label: 'Lists',
        to: '/home',
      },
      {
        id: 'profile',
        icon: `
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="28"
    height="28"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-user"
  >
    <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
    <circle cx="12" cy="7" r="4"></circle>
  </svg>`,
        label: 'Profile',
        to: '',
      },
      {
        id: 'more',
        icon: `
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="28"
    height="28"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="feather feather-more-horizontal"
  >
    <circle cx="12" cy="12" r="1"></circle>
    <circle cx="19" cy="12" r="1"></circle>
    <circle cx="5" cy="12" r="1"></circle>
  </svg>`,
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
        () => store.getters['userData'].handle,
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
    class="lg:w-1/5 border-r border-lighter dark:border-dark lg:px-8 py-2 flex flex-col justify-between"
  >
    <div v-show="ready">
      <button
        class="h-12 w-12 inline-flex items-center justify-center hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-20 text-blue rounded-full transition-colors duration-75"
      >
        <IconTwitter :size="32" />
      </button>
      <div>
        <router-link v-for="tab in tabs" :key="tab.id" :to="tab.to">
          <button
            @click="selectedTab = tab.id"
            class="flex items-center mr-auto mb-3 py-2 px-4 rounded-full hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-20 hover:text-blue focus:outline-none transition-colors duration-75"
            :class="
              selectedTab === tab.id
                ? 'text-blue dark:text-blue'
                : 'dark:text-lighter'
            "
          >
            <div class="mr-4" v-html="tab.icon"></div>
            <p class="text-left text-lg font-bold hidden lg:block">
              {{ tab.label }}
            </p>
          </button>
        </router-link>
        <button
          class="text-lightest bg-blue rounded-full font-semibold focus:outline-none w-12 h-12 lg:w-full lg:h-auto p-3 hover:bg-darkblue transition-colors duration-75"
          @click="showCreateFormDialog = true"
        >
          <p class="hidden lg:block">Tweet</p>
          <IconPlus
            :size="24"
            class="inline-flex items-center justify-center lg:hidden"
          />
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
          <p class="text-sm leading-tight dark:text-light">
            @{{ user.handle }}
          </p>
        </div>
        <IconChevronDown :size="18" class="ml-auto dark:text-lightest" />
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
              @{{ user.handle }}
            </p>
          </div>
          <IconCheck class="ml-auto text-blue" />
        </button>
        <button
          @click="logout"
          class="w-full text-left hover:bg-lightest dark:bg-black dark:hover:bg-darkest border-t border-lighter dark:border-dark p-3 text-sm dark:text-lightest focus:outline-none"
        >
          Log Out @{{ user.handle }}
        </button>
      </div>
    </div>
  </div>
</template>
