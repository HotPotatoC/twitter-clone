<script lang="ts">
import { defineComponent, ref, computed } from 'vue'
import IconImage from '../../icons/IconImage.vue'
import IconX from '../../icons/IconX.vue'

export default defineComponent({
  name: 'TweetCreateTweetForm',
  components: {
    IconImage,
    IconX,
  },
  emits: ['submit'],
  setup(_, { emit }) {
    const tweetContent = ref('')
    const tweetContentIsEmpty = computed(() => tweetContent.value === '')

    const tweetAttachmentInput = ref<HTMLInputElement>()
    const tweetAttachments = ref<File[]>([])
    const tweetAttachmentPreviews = ref<string[]>([])

    function onFileChange(event: Event) {
      const target = event.target as HTMLInputElement
      const file: File = target.files[0]

      if (file.size > 32 * 1024 * 1024) {
        tweetAttachmentInput.value.value = null
        return
      }

      if (tweetAttachments.value.length <= 4) {
        tweetAttachments.value.push(file)
        tweetAttachmentPreviews.value.push(URL.createObjectURL(file))

        // Set to null to be able to attach the same image before/after deletion
        tweetAttachmentInput.value.value = null
      } else {
        return
      }
    }

    function removeAttachment(index: number) {
      tweetAttachmentPreviews.value.splice(index, 1)
      tweetAttachments.value.splice(index, 1)
      tweetAttachmentInput.value.value = null
    }

    function submit() {
      emit('submit', {
        content: tweetContent.value,
        attachments: tweetAttachments.value,
      })
    }

    return {
      tweetContent,
      tweetContentIsEmpty,
      tweetAttachmentInput,
      tweetAttachments,
      tweetAttachmentPreviews,
      onFileChange,
      removeAttachment,
      submit,
    }
  },
})
</script>

<template>
  <form @submit.prevent="submit" class="w-full px-4 relative">
    <div
      @input="(e) => (tweetContent = e.target.innerText)"
      :contenteditable="true"
      placeholder="What's happening?"
      class="
        mt-3
        w-full
        focus:outline-none
        dark:bg-black
        dark:text-light
        whitespace-pre-wrap
        break-words
      "
    ></div>
    <div
      v-if="tweetAttachments.length > 0"
      class="relative overflow-hidden w-full"
    >
      <div
        v-if="tweetAttachments.length > 1"
        class="grid grid-cols-1 md:grid-cols-2 gap-4"
      >
        <div class="w-full">
          <div
            class="relative"
            :class="tweetAttachments.length > 2 ? 'h-48' : 'h-96'"
          >
            <button
              @click.prevent="removeAttachment(0)"
              class="
                absolute
                m-2
                p-1
                bg-black bg-opacity-75
                rounded-full
                focus:outline-none
              "
            >
              <IconX :size="24" class="text-white" />
            </button>
            <img
              v-lazy="tweetAttachmentPreviews[0]"
              class="object-cover w-full h-full rounded-lg"
            />
          </div>
          <div
            v-if="tweetAttachments.length > 2"
            class="relative"
            :class="tweetAttachments.length > 2 ? 'h-48' : 'h-96'"
          >
            <button
              @click.prevent="removeAttachment(1)"
              class="
                absolute
                m-2
                p-1
                bg-black bg-opacity-75
                rounded-full
                focus:outline-none
              "
            >
              <IconX :size="24" class="text-white" />
            </button>
            <img
              v-lazy="tweetAttachmentPreviews[1]"
              class="object-cover w-full h-full rounded-lg"
            />
          </div>
        </div>
        <div class="w-full">
          <div class="relative" v-if="tweetAttachments.length > 2">
            <button
              @click.prevent="removeAttachment(2)"
              class="
                absolute
                m-2
                p-1
                bg-black bg-opacity-75
                rounded-full
                focus:outline-none
              "
            >
              <IconX :size="24" class="text-white" />
            </button>
            <img
              v-lazy="tweetAttachmentPreviews[2]"
              class="object-cover w-full rounded-lg"
              :class="tweetAttachments.length === 4 ? 'h-full' : 'h-96'"
            />
          </div>
          <div
            class="relative"
            v-else
            :class="tweetAttachments.length > 2 ? 'h-48' : 'h-96'"
          >
            <button
              @click.prevent="removeAttachment(1)"
              class="
                absolute
                m-2
                p-1
                bg-black bg-opacity-75
                rounded-full
                focus:outline-none
              "
            >
              <IconX :size="24" class="text-white" />
            </button>
            <img
              v-lazy="tweetAttachmentPreviews[1]"
              class="object-cover w-full h-full rounded-lg"
            />
          </div>
          <div v-if="tweetAttachments.length === 4" class="relative h-48">
            <button
              @click.prevent="removeAttachment(3)"
              class="
                absolute
                m-2
                p-1
                bg-black bg-opacity-75
                rounded-full
                focus:outline-none
              "
            >
              <IconX :size="24" class="text-white" />
            </button>
            <img
              v-lazy="tweetAttachmentPreviews[3]"
              class="object-cover w-full h-full rounded-lg"
            />
          </div>
        </div>
      </div>
      <!-- Single image layout -->
      <div v-else class="w-full">
        <div class="relative">
          <button
            @click.prevent="removeAttachment(0)"
            class="
              absolute
              m-2
              p-1
              bg-black bg-opacity-75
              rounded-full
              focus:outline-none
            "
          >
            <IconX :size="24" class="text-white" />
          </button>
          <img
            v-lazy="tweetAttachmentPreviews[0]"
            class="object-cover w-full h-96 rounded-lg"
          />
        </div>
      </div>
    </div>
    <div class="my-6 border-b border-lighter dark:border-dark"></div>
    <button
      @click.prevent="tweetAttachmentInput.click()"
      class="p-2 rounded-full focus:outline-none"
      :class="
        tweetAttachments.length === 4
          ? 'cursor-default'
          : 'hover:bg-blue hover:bg-opacity-20'
      "
      :disabled="tweetAttachments.length === 4"
    >
      <IconImage
        :size="24"
        class="text-blue"
        :class="tweetAttachments.length === 4 ? 'opacity-50' : null"
      />
    </button>
    <input
      ref="tweetAttachmentInput"
      type="file"
      class="hidden"
      accept="image/jpg,image/jpeg,image/png"
      @change="onFileChange"
    />
    <button
      type="submit"
      class="
        h-10
        px-4
        font-semibold
        focus:outline-none
        rounded-full
        absolute
        bottom-0
        right-0
        transition-colors
        duration-75
      "
      :class="
        tweetContentIsEmpty
          ? ['bg-dark', 'text-light', 'cursor-default']
          : ['bg-blue', 'hover:bg-darkblue', 'text-lightest']
      "
      :disabled="tweetContentIsEmpty"
    >
      Tweet
    </button>
  </form>
</template>
