import { ActionTree } from 'vuex'
import { AugmentedActionContext } from '../../../store/types'
import { TweetsJSONSchema } from '../../tweets/store/actions'
import { Mutations, MutationTypes } from './mutations'
import { ProfileDetails, State } from './state'

import axios from '../../../services/axios'
import { Tweet } from '../../tweets/store/state'

export enum ActionTypes {
  REGISTER_ACCOUNT = 'REGISTER_ACCOUNT',
  GET_PROFILE_DETAILS = 'GET_PROFILE_DETAILS',
  GET_PROFILE_TWEETS = 'GET_PROFILE_TWEETS',
  LOAD_MORE_PROFILE_TWEETS = 'LOAD_MORE_PROFILE_TWEETS',
}

export interface Actions {
  [ActionTypes.REGISTER_ACCOUNT](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { name: string; email: string; password: string }
  ): Promise<void>
  [ActionTypes.GET_PROFILE_DETAILS](
    { commit }: AugmentedActionContext<Mutations, State>,
    username: string | string[]
  ): Promise<void>
  [ActionTypes.GET_PROFILE_TWEETS](
    { commit }: AugmentedActionContext<Mutations, State>,
    username: string | string[]
  ): Promise<void>
  [ActionTypes.LOAD_MORE_PROFILE_TWEETS](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { username: string | string[]; cursor: string }
  ): Promise<void>
}

interface GetProfileDetailsJSON {
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
}

export const actions: ActionTree<State, State> & Actions = {
  async [ActionTypes.REGISTER_ACCOUNT](
    { commit },
    { name, email, password }
  ): Promise<void> {
    try {
      await axios.post('/users/register', { name, email, password })
    } catch (error) {
      return error
    }
  },
  async [ActionTypes.GET_PROFILE_DETAILS]({ commit }, username): Promise<void> {
    try {
      const profileDetailsResponse = await axios.get<GetProfileDetailsJSON>(
        `/users/${username}`
      )

      const profileTweetsResponse = await axios.get<TweetsJSONSchema>(
        `/users/${username}/tweets`
      )

      const profile: ProfileDetails = {
        id: profileDetailsResponse.data.id,
        name: profileDetailsResponse.data.name,
        handle: profileDetailsResponse.data.handle,
        bio: profileDetailsResponse.data.bio,
        location: profileDetailsResponse.data.location,
        website: profileDetailsResponse.data.website,
        birthDate:
          profileDetailsResponse.data.birth_date !== '0001-01-01T00:00:00Z'
            ? profileDetailsResponse.data.birth_date
            : '',
        followersCount: profileDetailsResponse.data.followers_count,
        followingsCount: profileDetailsResponse.data.followings_count,
        joinedAt: profileDetailsResponse.data.joined_at,
      }

      commit(MutationTypes.SET_PROFILE_DETAILS, profile)
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: profileDetailsResponse.status,
        message: '',
      })
    } catch (error) {
      console.log(error)
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: error.response.status,
        message: error.response.data.message,
      })
    }
  },
  async [ActionTypes.GET_PROFILE_TWEETS]({ commit }, username): Promise<void> {
    try {
      const response = await axios.get<TweetsJSONSchema>(
        `/users/${username}/tweets`
      )

      const tweets: Tweet[] =
        response.data !== null
          ? response.data.items.map((item) => ({
              repliedToTweet: item.replied_to_tweet,
              repliedToName: item.replied_to_name,
              favoritesCount: item.favorites_count,
              repliesCount: item.replies_count,
              createdAt: item.created_at,
              ...item,
            }))
          : []

      commit(MutationTypes.SET_PROFILE_TWEETS, tweets)
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: response.status,
        message: '',
      })
    } catch (error) {
      console.log(error)
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: error.response.status,
        message: error.response.data.message,
      })
    }
  },
  async [ActionTypes.LOAD_MORE_PROFILE_TWEETS](
    { commit },
    { username, cursor }
  ): Promise<void> {
    try {
      const response = await axios.get<TweetsJSONSchema>(
        `/users/${username}/tweets?cursor=${cursor}`
      )

      const tweets: Tweet[] =
        response.data !== null
          ? response.data.items.map((item) => ({
              repliedToTweet: item.replied_to_tweet,
              repliedToName: item.replied_to_name,
              favoritesCount: item.favorites_count,
              repliesCount: item.replies_count,
              createdAt: item.created_at,
              ...item,
            }))
          : []

      commit(MutationTypes.PUSH_PROFILE_TWEETS, tweets)
    } catch (error) {
      console.log(error)
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: error.response.status,
        message: error.response.data.message,
      })
    }
  },
}
