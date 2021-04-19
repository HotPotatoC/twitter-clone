<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Root',
  setup() {},
})
</script>

<template>
  <div class="flex flex-wrap flex-col-reverse md:flex-row">
    <div
      class="w-full md:w-1/2 h-full relative flex justify-center items-center"
    >
      <img
        src="../../assets/twitter-banner.png"
        class="w-full h-screen object-cover"
        alt=""
      />
      <FontAwesome
        :icon="['fab', 'twitter']"
        class="__twitter-banner-icon-overlay absolute text-white"
      />
    </div>
    <div class="w-full md:w-1/2 h-screen px-6 md:px-12">
      <FontAwesome
        :icon="['fab', 'twitter']"
        class="mt-32 text-white text-5xl md:text-7xl"
      />
      <h1 class="mt-32 text-lighter text-5xl md:text-7xl font-bold">
        Happening now
      </h1>
      <h2 class="mt-12 text-lighter text-2xl md:text-4xl font-bold">
        Join Twitter today.
      </h2>

      <div class="flex flex-wrap flex-col space-y-6 mt-12">
        <button
          class="bg-blue text-lightest text-lg rounded-full font-semibold focus:outline-none h-auto p-4 hover:bg-darkblue transition-colors duration-75"
        >
          Sign Up
        </button>
        <router-link
          to="/login"
          class="border-2 border-blue text-lightest text-lg rounded-full font-semibold focus:outline-none h-auto p-4 hover:bg-blue hover:bg-opacity-10 transition-colors duration-75 text-center"
        >
          Log in
        </router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.__twitter-banner-icon-overlay {
  font-size: 12rem;
  line-height: 1;
}

@media (min-width: 768px) {
  .__twitter-banner-icon-overlay {
    font-size: 26rem;
    line-height: 1;
  }
}
</style>
