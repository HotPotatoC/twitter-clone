<script lang="ts">
import { defineComponent, toRefs } from 'vue'
import {
  TransitionRoot,
  TransitionChild,
  Dialog,
  DialogOverlay,
  DialogTitle,
  DialogDescription,
} from '@headlessui/vue'
import IconX from '../../icons/IconX.vue'

export default defineComponent({
  name: 'TweetImageOverlay',
  components: {
    TransitionRoot,
    TransitionChild,
    Dialog,
    DialogOverlay,
    DialogTitle,
    DialogDescription,
    IconX,
  },
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    image: {
      type: String,
      required: true,
    },
  },
  emits: ['close'],
  setup(props, { emit }) {
    const { show, image } = toRefs(props)

    function onClose() {
      emit('close')
    }

    return {
      show,
      onClose,
      image,
    }
  },
})
</script>

<template>
  <TransitionRoot appear :show="show" as="template">
    <Dialog as="div" static :open="show" @close="onClose">
      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div class="min-h-screen text-center">
          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0"
            enter-to="opacity-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100"
            leave-to="opacity-0"
          >
            <DialogOverlay class="fixed inset-0 bg-blue bg-opacity-20" />
          </TransitionChild>

          <span class="inline-block h-screen align-middle" aria-hidden="true">
            &#8203;
          </span>

          <TransitionChild
            as="template"
            enter="duration-100 ease-out"
            enter-from="opacity-0"
            enter-to="opacity-100"
            leave="duration-100 ease-in"
            leave-from="opacity-100"
            leave-to="opacity-0"
          >
            <div
              class="inline-block w-full h-full overflow-hidden align-middle transition-all transform"
            >
              <button
                @click="onClose()"
                class="absolute top-0 left-0 m-4 p-1 border-none rounded-full focus:outline-none hover:bg-white hover:bg-opacity-20"
              >
                <IconX :size="26" class="dark:text-white" />
              </button>
              <div class="flex items-center h-full">
                <div class="w-full flex justify-center">
                  <img v-lazy="image" class="max-w-5xl" />
                </div>
                <div
                  class="w-1/4 h-screen overflow-y-scroll bg-black border-l border-lighter dark:border-darker"
                ></div>
              </div>
            </div>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
