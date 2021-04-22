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
import NavigationSidebarTab from './NavigationSidebarTab.vue'
import IconHome from '../icons/IconHome.vue'
import IconHashtag from '../icons/IconHashtag.vue'
import IconBell from '../icons/IconBell.vue'
import IconEnvelope from '../icons/IconEnvelope.vue'
import IconBookmark from '../icons/IconBookmark.vue'
import IconList from '../icons/IconList.vue'
import IconUser from '../icons/IconUser.vue'
import IconEllipsisH from '../icons/IconEllipsisH.vue'

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
    NavigationSidebarTab,
    IconHome,
    IconHashtag,
    IconBell,
    IconEnvelope,
    IconBookmark,
    IconList,
    IconUser,
    IconEllipsisH,
  },
  setup() {
    const store = useStore()
    const router = useRouter()
    const selectedTab = ref<string>('home')
    const showDropdown = ref<boolean>(false)
    const showCreateFormDialog = ref<boolean>(false)

    const user = computed(() => store.getters['userData'])

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

    function selectTab(tab: string) {
      selectedTab.value = tab
    }

    return {
      selectedTab,
      selectTab,
      showDropdown,
      showCreateFormDialog,
      user,
      logout,
      createTweet,
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
    <div>
      <button
        class="h-12 w-12 inline-flex items-center justify-center hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-20 text-blue rounded-full transition-colors duration-75"
      >
        <IconTwitter :size="32" />
      </button>
      <NavigationSidebarTab
        id="home"
        label="Home"
        to="/home"
        :selected="selectedTab === 'home'"
        @selectTab="selectTab"
      >
        <template #icon>
          <IconHome :size="28" class="mr-4" />
        </template>
      </NavigationSidebarTab>
      <NavigationSidebarTab
        id="explore"
        label="Explore"
        to="/home"
        :selected="selectedTab === 'explore'"
        @selectTab="selectTab"
      >
        <template #icon>
          <IconHashtag :size="28" class="mr-4" />
        </template>
      </NavigationSidebarTab>
      <NavigationSidebarTab
        id="notifications"
        label="Notifications"
        to="/home"
        :selected="selectedTab === 'notifications'"
        @selectTab="selectTab"
      >
        <template #icon>
          <IconBell :size="28" class="mr-4" />
        </template>
      </NavigationSidebarTab>
      <NavigationSidebarTab
        id="messages"
        label="Messages"
        to="/home"
        :selected="selectedTab === 'messages'"
        @selectTab="selectTab"
      >
        <template #icon>
          <IconEnvelope :size="28" class="mr-4" />
        </template>
      </NavigationSidebarTab>
      <NavigationSidebarTab
        id="bookmarks"
        label="Bookmarks"
        to="/home"
        :selected="selectedTab === 'bookmarks'"
        @selectTab="selectTab"
      >
        <template #icon>
          <IconBookmark :size="28" class="mr-4" />
        </template>
      </NavigationSidebarTab>
      <NavigationSidebarTab
        id="lists"
        label="Lists"
        to="/home"
        :selected="selectedTab === 'lists'"
        @selectTab="selectTab"
      >
        <template #icon>
          <IconList :size="28" class="mr-4" />
        </template>
      </NavigationSidebarTab>
      <NavigationSidebarTab
        id="profile"
        label="Profile"
        :to="`/${user.handle}`"
        :selected="selectedTab === 'profile'"
        @selectTab="selectTab"
      >
        <template #icon>
          <IconUser :size="28" class="mr-4" />
        </template>
      </NavigationSidebarTab>
      <NavigationSidebarTab
        id="more"
        label="More"
        to="/home"
        :selected="selectedTab === 'more'"
        @selectTab="selectTab"
      >
        <template #icon>
          <IconEllipsisH :size="28" class="mr-4" />
        </template>
      </NavigationSidebarTab>
    </div>
    <div class="lg:w-full relative">
      <button
        @click="showDropdown = !showDropdown"
        class="flex items-center w-full hover:bg-lightblue dark:hover:bg-darkblue dark:hover:bg-opacity-10 rounded-full p-2 focus:outline-none transition-colors duration-75"
      >
        <div class="lg:ml-4">
          <span class="text-sm font-bold leading-tight dark:text-lightest">
            {{ user.name }}
          </span>
          <span class="text-sm leading-tight dark:text-light">
            @{{ user.handle }}
          </span>
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
            <span class="text-sm font-bold leading-tight dark:text-lightest">
              {{ user.name }}
            </span>
            <span class="text-sm leading-tight dark:text-light">
              @{{ user.handle }}
            </span>
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
