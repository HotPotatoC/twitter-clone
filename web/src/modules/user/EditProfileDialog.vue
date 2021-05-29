<script lang="ts">
import {
  defineComponent,
  toRefs,
  ref,
  reactive,
  computed,
  watchEffect,
} from 'vue'
import { Cropper, CircleStencil } from 'vue-advanced-cropper'
import 'vue-advanced-cropper/dist/style.css'
import { ProfileDetails } from './types'
import * as utils from '../../utils'
import IconX from '../../icons/IconX.vue'
import Dialog from '../../shared/Dialog.vue'
import { UpdatableProfileFieldsReactive } from './types'
import { Month, Birthdate } from '../../types'
import IconCamera from '../../icons/IconCamera.vue'
import { useStore } from '../../store'
import { Action } from '../storeActionTypes'
import { useRoute } from 'vue-router'

export default defineComponent({
  name: 'EditProfileDialog',
  components: {
    Dialog,
    IconX,
    IconCamera,
    Cropper,
    CircleStencil,
  },
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    profile: {
      type: Object as () => ProfileDetails,
      required: true,
    },
  },
  setup(props, { emit }) {
    const { show, profile } = toRefs(props)
    const store = useStore()
    const route = useRoute()

    const name = ref(profile.value.name)
    const bio = ref(profile.value.bio)
    const location = ref(profile.value.location)
    const website = ref(profile.value.website)

    const selectedYear = ref(new Date(profile.value.birthDate).getFullYear())
    const selectedMonth = ref(new Date(profile.value.birthDate).getMonth() - 1)
    const selectedDay = ref(new Date(profile.value.birthDate).getDate())

    const birthDate = ref<Birthdate>(
      `${selectedYear.value}-${(selectedMonth.value + 1)
        .toString()
        .padStart(2, '0')}-${selectedDay.value
        .toString()
        .padStart(2, '0')}` as Birthdate
    )

    const availableDays = ref(
      utils.getDaysInMonth(selectedMonth.value, selectedYear.value)
    )

    const editProfileData = reactive<UpdatableProfileFieldsReactive>({
      name,
      bio,
      location,
      website,
      birthDate,
    })

    const months = computed<Month[]>(() => [
      'January',
      'February',
      'March',
      'April',
      'May',
      'June',
      'July',
      'August',
      'September',
      'October',
      'November',
      'December',
    ])

    const years = computed(() => utils.getYears())

    watchEffect(() => {
      availableDays.value = utils.getDaysInMonth(
        selectedMonth.value,
        selectedYear.value
      )
    })

    watchEffect(() => {
      birthDate.value = `${selectedYear.value}-${(selectedMonth.value + 1)
        .toString()
        .padStart(2, '0')}-${selectedDay.value
        .toString()
        .padStart(2, '0')}` as Birthdate
    })

    const cropper = ref()
    const showCropper = ref(false)
    const profileImageInput = ref<HTMLInputElement>()
    const profileImageSrc = ref('')
    const profileImageName = ref('')
    const profileImageType = ref('')

    function onFileChange(event: Event) {
      const target = event.target as HTMLInputElement
      const file: File = target.files[0]

      if (file.size > 2.5 * 1024 * 1024) {
        profileImageInput.value.value = null
        return
      }

      showCropper.value = true
      profileImageSrc.value = URL.createObjectURL(file)
      profileImageName.value = file.name

      const reader = new FileReader()

      reader.onload = (e) => {
        profileImageType.value = utils.getMimeType(e.target.result, file.type)
      }

      reader.readAsArrayBuffer(file)
    }

    async function updateProfileImage() {
      const canvas = cropper.value.getResult().canvas as HTMLCanvasElement
      if (canvas) {
        const blob = await new Promise<Blob>((resolve) =>
          canvas.toBlob(resolve, profileImageType.value)
        )

        try {
          await store.dispatch(Action.UserActionTypes.UPDATE_PROFILE_IMAGE, {
            image: blob,
            fileName: profileImageName.value,
          })

          await store.dispatch(
            Action.UserActionTypes.GET_PROFILE_DETAILS,
            route.params.name as string
          )
          close()
        } catch (error) {
          console.error(error)
        }
      }
    }

    function close() {
      showCropper.value = false
      profileImageInput.value.value = null
      URL.revokeObjectURL(profileImageSrc.value)
      emit('close')
    }

    function dispatch() {
      emit('dispatch', editProfileData)
    }

    return {
      CircleStencil,
      show,
      close,
      dispatch,
      profile,
      name,
      bio,
      location,
      website,
      months,
      years,
      availableDays,
      selectedYear,
      selectedMonth,
      selectedDay,
      profileImageInput,
      profileImageSrc,
      showCropper,
      onFileChange,
      cropper,
      updateProfileImage,
    }
  },
})
</script>

<template>
  <Dialog size="2xl" :show="show" @close="close" :closeButton="false">
    <template #title>
      <div class="flex justify-between items-center pb-4">
        <div class="flex">
          <button
            @click="close"
            class="text-lg font-medium leading-6 text-blue focus:outline-none"
          >
            <IconX :size="24" />
          </button>
          <h1 class="ml-6 font-bold text-2xl dark:text-lightest">
            Edit Profile
          </h1>
        </div>
        <button
          v-if="!showCropper"
          @click="dispatch"
          class="
            h-10
            px-4
            text-lightest
            bg-blue
            font-semibold
            focus:outline-none
            rounded-full
            transition-colors
            duration-75
          "
        >
          Save
        </button>
        <button
          v-else
          @click="updateProfileImage"
          class="
            h-10
            px-4
            text-lightest
            bg-blue
            font-semibold
            focus:outline-none
            rounded-full
            transition-colors
            duration-75
          "
        >
          Apply
        </button>
      </div>
    </template>
    <div v-if="!showCropper">
      <div class="w-full px-2 h-48 mb-16 bg-blue relative">
        <button
          @click.prevent="profileImageInput.click()"
          class="
            absolute
            mt-32
            overflow-hidden
            rounded-full
            w-32
            h-32
            z-10
            bg-dark bg-opacity-30
            focus:outline-none
            flex
            items-center
            justify-center
          "
        >
          <IconCamera :size="24" class="text-white opacity-75" />
        </button>
        <input
          ref="profileImageInput"
          type="file"
          class="hidden"
          accept="image/jpg,image/jpeg,image/png"
          @change="onFileChange"
        />
        <div
          class="
            absolute
            overflow-hidden
            mt-32
            rounded-full
            w-32
            h-32
            border-4 border-lightest
            dark:border-black
          "
        >
          <img v-lazy="profile.photoURL" class="w-32 h-32" />
        </div>
      </div>
      <input
        v-model="name"
        type="text"
        placeholder="Name"
        class="
          w-full
          px-2
          py-4
          my-4
          border-2 border-lighter
          text-xl
          rounded
          dark:border-dark
          focus:outline-none
          dark:bg-black
          dark:text-light
          focus:border-blue
          dark:focus:border-blue
          transition-colors
          duration-75
        "
      />
      <textarea
        v-model="bio"
        placeholder="Bio"
        class="
          w-full
          px-2
          py-4
          my-4
          border-2 border-lighter
          text-xl
          rounded
          dark:border-dark
          focus:outline-none
          dark:bg-black
          dark:text-light
          focus:border-blue
          dark:focus:border-blue
          transition-colors
          duration-75
        "
      />
      <input
        v-model="location"
        type="text"
        placeholder="Location"
        class="
          w-full
          px-2
          py-4
          my-4
          border-2 border-lighter
          text-xl
          rounded
          dark:border-dark
          focus:outline-none
          dark:bg-black
          dark:text-light
          focus:border-blue
          dark:focus:border-blue
          transition-colors
          duration-75
        "
      />
      <input
        v-model="website"
        type="text"
        placeholder="Website"
        class="
          w-full
          px-2
          py-4
          my-4
          border-2 border-lighter
          text-xl
          rounded
          dark:border-dark
          focus:outline-none
          dark:bg-black
          dark:text-light
          focus:border-blue
          dark:focus:border-blue
          transition-colors
          duration-75
        "
      />
      <div class="mt-6">
        <span class="font-bold text-2xl dark:text-lightest">Birth date</span>
      </div>
      <div class="flex space-x-4">
        <select
          v-model="selectedYear"
          class="
            relative
            appearance-none
            w-full
            px-2
            py-4
            my-4
            border-2 border-lighter
            text-xl
            rounded
            dark:border-dark
            focus:outline-none
            dark:bg-black
            dark:text-light
            focus:border-blue
            dark:focus:border-blue
            cursor-pointer
            transition-colors
            duration-75
          "
        >
          <option v-for="(year, index) in years" :key="index" :value="year">
            {{ year }}
          </option>
        </select>
        <select
          v-model="selectedMonth"
          class="
            relative
            appearance-none
            w-full
            px-2
            py-4
            my-4
            border-2 border-lighter
            text-xl
            rounded
            dark:border-dark
            focus:outline-none
            dark:bg-black
            dark:text-light
            focus:border-blue
            dark:focus:border-blue
            cursor-pointer
            transition-colors
            duration-75
          "
        >
          <option v-for="(month, index) in months" :key="index" :value="index">
            {{ month }}
          </option>
        </select>
        <select
          v-model="selectedDay"
          class="
            relative
            appearance-none
            w-full
            px-2
            py-4
            my-4
            border-2 border-lighter
            text-xl
            rounded
            dark:border-dark
            focus:outline-none
            dark:bg-black
            dark:text-light
            focus:border-blue
            dark:focus:border-blue
            cursor-pointer
            transition-colors
            duration-75
          "
        >
          <option
            v-for="(day, index) in availableDays"
            :key="index"
            :value="day.getDate()"
          >
            {{ day.getDate() }}
          </option>
        </select>
      </div>
    </div>
    <div v-else>
      <Cropper
        class="cropper"
        ref="cropper"
        :src="profileImageSrc"
        :stencil-props="{
          aspectRatio: 1 / 1,
        }"
        :auto-zoom="true"
        :stencil-component="CircleStencil"
      />
    </div>
  </Dialog>
</template>
