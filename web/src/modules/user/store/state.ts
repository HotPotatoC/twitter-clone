import { Tweet } from '../../tweets/store/state'

export interface ProfileStatus {
  statusCode: number
  message: string
}

export interface ProfileDetails {
  id: number
  name: string
  bio: string
  location: string
  website: string
  birthDate: string
  followersCount: number
  followingsCount: number
  joinedAt: string
}

export interface State {
  status?: ProfileStatus
  profileDetails: ProfileDetails
  profileTweets: Tweet[]
}

export const state: State = {
  profileDetails: {
    id: 0,
    name: '',
    bio: '',
    location: '',
    website: '',
    birthDate: '',
    followersCount: 0,
    followingsCount: 0,
    joinedAt: '',
  },
  profileTweets: [],
}
