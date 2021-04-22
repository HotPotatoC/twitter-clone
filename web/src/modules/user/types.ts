export interface ProfileDetailsJSONSchema {
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

export interface ProfileStatus {
  statusCode: number
  message: string
}

export interface ProfileDetails {
  id: number
  name: string
  handle: string
  bio: string
  location: string
  website: string
  birthDate: string
  followersCount: number
  followingsCount: number
  joinedAt: string
}
