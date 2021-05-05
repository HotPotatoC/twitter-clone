<script lang="ts">
import { defineComponent, toRefs } from 'vue'
import Dialog from '../../shared/Dialog.vue'
import TweetCreateTweetForm from './TweetCreateTweetForm.vue'

export default defineComponent({
  name: 'TweetCreateTweetDialog',
  components: { Dialog, TweetCreateTweetForm },
  props: {
    show: { type: Boolean, default: false },
  },
  emits: ['submit', 'close'],
  setup(props, { emit }) {
    const { show } = toRefs(props)

    function submit(payload) {
      emit('submit', payload)
    }

    function close() {
      emit('close')
    }

    return {
      show,
      emit,
      submit,
      close,
    }
  },
})
</script>

<template>
  <Dialog size="2xl" :show="show" @close="close">
    <TweetCreateTweetForm @submit="submit" />
  </Dialog>
</template>
