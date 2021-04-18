<template>
  <div
    class="md:block hidden w-1/2 h-full border-l border-lighter dark:border-darker py-2 px-6 overflow-y-scroll relative"
  >
    <form v-show="!isInSearchPage" @submit.prevent="redirectWithSearchQuery">
      <input
        class="pl-12 rounded-full w-full p-3 bg-lighter dark:bg-darkest dark:text-light text-sm mb-4 focus:bg-black focus:outline-none border-2 border-lighter dark:border-darkest focus:border-blue dark:focus:text-lightest transition duration-150"
        @focus="searchFocused = true"
        @blur="searchFocused = false"
        v-model="searchQuery"
        type="search"
        placeholder="Search Twitter"
      />
      <input type="submit" class="hidden" />
      <FontAwesome
        :icon="['fas', 'search']"
        class="absolute left-0 top-0 mt-6 ml-12 text-base"
        :class="searchFocused ? 'text-blue' : 'text-light'"
      />
    </form>
    <div
      class="w-full rounded-2xl bg-lightest dark:bg-darkest my-4 hidden lg:block"
    >
      <div class="p-4">
        <p class="text-lg font-bold dark:text-lightest">Who to Follow</p>
      </div>
      <button
        v-for="friend in friends"
        class="w-full flex hover:bg-lighter dark:hover:bg-darker p-3 border-t border-lighter dark:border-darker focus:outline-none"
      >
        <div>
          <p class="text-sm font-bold leading-tight dark:text-lightest">
            {{ friend.name }}
          </p>
          <p class="text-sm leading-tight dark:text-light">
            {{ friend.handle }}
          </p>
        </div>
        <button
          class="ml-auto text-sm text-blue py-1 px-4 rounded-full border-2 border-blue hover:bg-blue hover:bg-opacity-25 focus:outline-none"
        >
          Follow
        </button>
      </button>
      <button
        class="p-4 w-full hover:bg-lighter dark:hover:bg-darker text-left text-blue border-t border-lighter dark:border-darker rounded-b-2xl focus:outline-none"
      >
        Show More
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

export default defineComponent({
  name: 'TrendingSidebar',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const searchFocused = ref<boolean>(false)
    const searchQuery = ref<string>('')

    const isInSearchPage = computed(() => route.path === '/search')

    function redirectWithSearchQuery() {
      router.push({
        path: '/search',
        query: { q: searchQuery.value },
      })
      return
    }

    return {
      searchFocused,
      searchQuery,
      redirectWithSearchQuery,
      isInSearchPage,
      friends: [
        { name: 'Elon Musk', handle: '@teslaBoy' },
        { name: 'Adrian Monk', handle: '@detective' },
        { name: 'Kevin Hart', handle: '@miniRock' },
      ],
    }
  },
})
</script>
