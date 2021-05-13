<script lang="ts">
import { defineComponent, toRefs } from 'vue'
import {
  TransitionRoot,
  TransitionChild,
  Dialog as BaseDialog,
  DialogOverlay as BaseDialogOverlay,
  DialogTitle as BaseDialogTitle,
  DialogDescription as BaseDialogDescription,
} from '@headlessui/vue'
import IconX from '../icons/IconX.vue'

export default defineComponent({
  name: 'Dialog',
  components: {
    TransitionRoot,
    TransitionChild,
    BaseDialog,
    BaseDialogOverlay,
    BaseDialogTitle,
    BaseDialogDescription,
    IconX,
  },
  props: {
    size: {
      type: String as () =>
        | 'xs'
        | 'sm'
        | 'md'
        | 'lg'
        | 'xl'
        | '2xl'
        | '3xl'
        | '4xl'
        | '5xl'
        | '6xl'
        | '7xl'
        | 'full',
      default: 'md',
    },
    closeButton: {
      type: Boolean,
      default: true,
    },
    show: {
      type: Boolean,
      default: false,
    },
  },
  emits: ['close'],
  setup(props, { emit }) {
    const { show } = toRefs(props)

    function onClose() {
      emit('close')
    }

    return {
      show,
      onClose,
    }
  },
})
</script>

<template>
  <TransitionRoot appear :show="show" as="template">
    <BaseDialog as="div" static :open="show" @close="onClose">
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
            <BaseDialogOverlay
              class="fixed inset-0 bg-darkblue bg-opacity-20"
            />
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
              class="
                inline-block
                w-full
                p-6
                my-8
                overflow-hidden
                text-left
                align-middle
                transition-all
                transform
                bg-white
                dark:bg-black
                shadow-xl
                rounded-2xl
              "
              :class="`max-w-${size}`"
            >
              <BaseDialogTitle
                v-if="closeButton"
                as="button"
                @click="onClose()"
                class="
                  text-lg
                  font-medium
                  leading-6
                  text-blue
                  focus:outline-none
                "
              >
                <IconX :size="24" />
              </BaseDialogTitle>
              <BaseDialogTitle v-else>
                <slot name="title"></slot>
              </BaseDialogTitle>
              <div
                :class="
                  closeButton
                    ? ['mt-2', 'border-t', 'border-lighter', 'dark:border-dark']
                    : null
                "
              >
                <slot></slot>
              </div>
            </div>
          </TransitionChild>
        </div>
      </div>
    </BaseDialog>
  </TransitionRoot>
</template>
