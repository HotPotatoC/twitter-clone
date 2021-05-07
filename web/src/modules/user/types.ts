import { Ref } from 'vue'

export type ProfileDetailsJSONSchema = {
  id: number
  name: string
  handle: string
  bio: string
  location: string
  website: string
  birth_date: string
  photo_url: string
  followers_count: number
  followings_count: number
  is_following: boolean
  joined_at: string
}

export type ProfileStatus = {
  statusCode: number
  message: string
}

export type ProfileDetails = {
  id: number
  name: string
  handle: string
  bio: string
  location: string
  website: string
  birthDate: string
  photoURL: string
  followersCount: number
  followingsCount: number
  isFollowing: boolean
  joinedAt: string
}

export type UpdatableProfileFields = Partial<
  Pick<
    ProfileDetails,
    'name' | 'bio' | 'location' | 'website' | 'birthDate' | 'photoURL'
  >
>

export type UpdatableProfileFieldsReactive = Partial<
  Omit<
    UpdatableProfileFields,
    'name' | 'bio' | 'location' | 'website' | 'birthDate' | 'photoURL'
  > & {
    name: string | Ref<string>
    bio: string | Ref<string>
    location: string | Ref<string>
    website: string | Ref<string>
    birthDate: string | Ref<string>
  }
>
