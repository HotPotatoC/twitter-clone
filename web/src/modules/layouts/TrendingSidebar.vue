<script lang="ts">
import { computed, defineComponent, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import IconSearch from '../../icons/IconSearch.vue'

export default defineComponent({
  components: { IconSearch },
  name: 'TrendingSidebar',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const searchFocused = ref(false)
    const searchQuery = ref('')

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

<template>
  <div
    class="
      md:block
      hidden
      w-1/2
      h-full
      border-l border-lighter
      dark:border-dark
      py-2
      px-6
      overflow-y-scroll
      relative
    "
  >
    <form v-show="!isInSearchPage" @submit.prevent="redirectWithSearchQuery">
      <input
        class="
          pl-12
          rounded-full
          w-full
          p-3
          bg-lighter
          dark:bg-darkest
          dark:text-light
          text-sm
          mb-4
          focus:bg-white
          dark:focus:bg-black
          focus:outline-none
          border-2 border-lighter
          dark:border-darkest
          focus:border-blue
          dark:focus:border-blue
          dark:focus:text-lightest
          transition
          duration-150
        "
        @focus="searchFocused = true"
        @blur="searchFocused = false"
        v-model="searchQuery"
        type="search"
        placeholder="Search Twitter"
      />
      <IconSearch
        :size="24"
        class="absolute left-0 top-0 mt-5 ml-10"
        :class="searchFocused ? 'text-blue' : 'text-light'"
      />
      <input type="submit" class="hidden" />
    </form>
    <div
      class="
        w-full
        rounded-2xl
        bg-lightest
        dark:bg-darkest
        my-4
        hidden
        lg:block
      "
    >
      <div class="p-4">
        <p class="text-lg font-bold dark:text-lightest">Who to Follow</p>
      </div>
      <button
        v-for="friend in friends"
        class="
          w-full
          flex
          hover:bg-lighter
          dark:hover:bg-darker
          p-3
          border-t border-lighter
          dark:border-dark
          focus:outline-none
          transition-colors
          duration-75
        "
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
          class="
            ml-auto
            text-sm text-blue
            py-1
            px-4
            rounded-full
            border-2 border-blue
            hover:bg-blue hover:bg-opacity-25
            focus:outline-none
            transition-colors
            duration-75
          "
        >
          Follow
        </button>
      </button>
      <button
        class="
          p-4
          w-full
          hover:bg-lighter
          dark:hover:bg-darker
          text-left text-blue
          border-t border-lighter
          dark:border-dark
          rounded-b-2xl
          focus:outline-none
          transition-colors
          duration-75
        "
      >
        Show More
      </button>
    </div>
  </div>
</template>
