<script lang="ts">
import { defineComponent, toRefs, ref, Ref, reactive, computed } from 'vue'
import Dialog from '../../shared/Dialog.vue'

interface NewReply {
  content: string | Ref<string>
}

export default defineComponent({
  name: 'TweetCreateReplyDialog',
  components: { Dialog },
  props: {
    show: {
      type: Boolean,
      default: false,
    },
  },
  emit: ['close', 'dispatch'],
  setup(props, { emit }) {
    const { show } = toRefs(props)
    const replyContent = ref('')
    const newReply = reactive<NewReply>({
      content: replyContent,
    })

    const contentIsEmpty = computed(() => replyContent.value === '')

    function close() {
      emit('close')
    }

    function dispatch() {
      emit('dispatch', newReply.content)
    }

    return { show, emit, newReply, contentIsEmpty, close, dispatch }
  },
})
</script>

<template>
  <Dialog :show="show" @close="close">
    <form @submit.prevent="dispatch" class="w-full">
      <textarea
        v-model="newReply.content"
        placeholder="Tweet your reply"
        class="mt-3 w-full focus:outline-none dark:bg-black dark:text-light"
      />

      <div class="mt-4 text-right">
        <button
          type="submit"
          class="
            h-10
            px-4
            font-semibold
            focus:outline-none
            rounded-full
            transition-colors
            duration-75
          "
          :class="
            contentIsEmpty
              ? ['bg-dark', 'text-light', 'cursor-default']
              : ['bg-blue', 'hover:bg-darkblue', 'text-lightest']
          "
          @click="close"
          :disabled="contentIsEmpty"
        >
          Tweet
        </button>
      </div>
    </form>
  </Dialog>
</template>
