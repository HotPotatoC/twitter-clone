<template>
  <TransitionRoot appear :show="show" as="template">
    <Dialog as="div" static :open="show" @close="onClose">
      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="min-h-screen px-4 text-center">
          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0"
            enter-to="opacity-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100"
            leave-to="opacity-0"
          >
            <DialogOverlay class="fixed inset-0 bg-darkblue bg-opacity-20" />
          </TransitionChild>

          <span class="inline-block h-screen align-middle" aria-hidden="true">
            &#8203;
          </span>

          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
            <div
              class="inline-block w-full max-w-md p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-white dark:bg-black shadow-xl rounded-2xl"
            >
              <DialogTitle
                as="button"
                @click="onClose()"
                class="text-lg font-medium leading-6 text-blue focus:outline-none"
              >
                <FontAwesome :icon="['fas', 'times']" />
              </DialogTitle>
              <div
                class="mt-2 border-t border-lighter dark:border-light dark:border-opacity-25"
              >
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
                      @click="onClose()"
                      :disabled="contentIsEmpty"
                    >
                      Tweet
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>

<script lang="ts">
import { defineComponent, ref, Ref, reactive, toRefs, computed } from 'vue'
import { useStore } from '../../store'
import {
  TransitionRoot,
  TransitionChild,
  Dialog,
  DialogOverlay,
  DialogTitle,
  DialogDescription,
} from '@headlessui/vue'
import { ActionTypes } from '../../modules/tweets/store/actions'

interface NewTweet {
  content: string | Ref<string>
}

export default defineComponent({
  name: 'TweetCreateFormDialog',
  components: {
    TransitionRoot,
    TransitionChild,
    Dialog,
    DialogOverlay,
    DialogTitle,
    DialogDescription,
  },
  props: {
    show: {
      type: Boolean,
      default: false,
    },
  },
  emits: ['close'],
  setup(props, { emit }) {
    const { show } = toRefs(props)
    const store = useStore()

    const tweetContent = ref<string>('')
    const newTweet = reactive<NewTweet>({
      content: tweetContent,
    })

    const contentIsEmpty = computed(() => tweetContent.value === '')

    function onClose() {
      emit('close')
    }

    async function addNewTweet() {
      try {
        await store.dispatch(ActionTypes.NEW_TWEET, newTweet.content)
        newTweet.content = ''
      } catch (error) {
        console.log(error)
      }
    }

    return {
      show,
      onClose,
      contentIsEmpty,
      newTweet,
      addNewTweet,
    }
  },
})
</script>
