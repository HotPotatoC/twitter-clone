import { Ref } from 'vue'
import { Birthdate } from '../../types'

export type ProfileDetailsJSONSchema = {
  id: number
  name: string
  handle: string
  bio: string
  location: string
  website: string
  birth_date: string
  followers_count: number
  followings_count: number
  joined_at: string
  already_liked: boolean
}

export type ProfileStatus = {
  statusCode: number
  message: string
}

export type ProfileDescription = {
  name: string
  bio: string
  location: string
  website: string
  birthDate: string
}

export type ProfileDetails = {
  id: number
  handle: string
  followersCount: number
  followingsCount: number
  joinedAt: string
} & ProfileDescription

export type EditProfilePayload = {
  name?: string | Ref<string>
  bio?: string | Ref<string>
  location?: string | Ref<string>
  website?: string | Ref<string>
  birthDate?: Birthdate | Ref<Birthdate>
}
