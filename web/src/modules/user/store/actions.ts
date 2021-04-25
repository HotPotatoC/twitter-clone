import { ActionTree } from 'vuex'
import { Mutations, MutationTypes } from './mutations'
import { State } from './state'
import {
  fetchProfileDetails,
  followUser,
  registerAccount,
  unfollowUser,
  updateProfile,
} from '../service'
import { fetchUserTweets } from '../../tweets/service'
import { AugmentedActionContext } from '../../../store'
import { ProfileDescription } from '../types'

export enum ActionTypes {
  REGISTER_ACCOUNT = 'REGISTER_ACCOUNT',
  GET_PROFILE_DETAILS = 'GET_PROFILE_DETAILS',
  GET_PROFILE_TWEETS = 'GET_PROFILE_TWEETS',
  LOAD_MORE_PROFILE_TWEETS = 'LOAD_MORE_PROFILE_TWEETS',
  UPDATE_PROFILE_DETAILS = 'UPDATE_PROFILE_DETAILS',
  FOLLOW_USER = 'FOLLOW_USER',
  UNFOLLOW_USER = 'UNFOLLOW_USER',
}

export type Actions = {
  [ActionTypes.REGISTER_ACCOUNT](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { name: string; email: string; password: string }
  ): Promise<void>
  [ActionTypes.GET_PROFILE_DETAILS](
    { commit }: AugmentedActionContext<Mutations, State>,
    handle: string
  ): Promise<void>
  [ActionTypes.GET_PROFILE_TWEETS](
    { commit }: AugmentedActionContext<Mutations, State>,
    handle: string
  ): Promise<void>
  [ActionTypes.LOAD_MORE_PROFILE_TWEETS](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { handle: string; cursor: string }
  ): Promise<void>
  [ActionTypes.UPDATE_PROFILE_DETAILS](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: ProfileDescription
  ): Promise<void>
  [ActionTypes.FOLLOW_USER](
    { commit }: AugmentedActionContext<Mutations, State>,
    userId: string
  ): Promise<void>
  [ActionTypes.UNFOLLOW_USER](
    { commit }: AugmentedActionContext<Mutations, State>,
    userId: string
  ): Promise<void>
}

export const actions: ActionTree<State, State> & Actions = {
  async [ActionTypes.REGISTER_ACCOUNT](
    { commit },
    { name, email, password }
  ): Promise<void> {
    try {
      await registerAccount({ name, email, password })
    } catch (error) {
      return error
    }
  },
  async [ActionTypes.GET_PROFILE_DETAILS]({ commit }, handle): Promise<void> {
    try {
      const profile = await fetchProfileDetails(handle)

      commit(MutationTypes.SET_PROFILE_DETAILS, profile)
    } catch (error) {
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: error.response.status,
        message: error.response.data.message,
      })
    }
  },
  async [ActionTypes.GET_PROFILE_TWEETS]({ commit }, handle): Promise<void> {
    try {
      const tweets = await fetchUserTweets(handle)

      commit(MutationTypes.SET_PROFILE_TWEETS, tweets)
    } catch (error) {
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: error.response.status,
        message: error.response.data.message,
      })
    }
  },
  async [ActionTypes.LOAD_MORE_PROFILE_TWEETS](
    { commit },
    { handle, cursor }
  ): Promise<void> {
    try {
      const tweets = await fetchUserTweets(handle, cursor)

      commit(MutationTypes.PUSH_PROFILE_TWEETS, tweets)
    } catch (error) {
      commit(MutationTypes.SET_PROFILE_STATUS, {
        statusCode: error.response.status,
        message: error.response.data.message,
      })
    }
  },
  async [ActionTypes.UPDATE_PROFILE_DETAILS](
    { commit },
    { name, bio, location, website, birthDate }
  ): Promise<void> {
    try {
      await updateProfile({
        name,
        bio,
        location,
        website,
        birthDate,
      })

      commit(MutationTypes.UPDATE_PROFILE, {
        name,
        bio,
        location,
        website,
        birthDate,
      })
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.FOLLOW_USER]({ commit }, userId): Promise<void> {
    try {
      await followUser(userId)

      commit(MutationTypes.SET_IS_FOLLOWING_USER, true)
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.UNFOLLOW_USER]({ commit }, userId): Promise<void> {
    try {
      await unfollowUser(userId)

      commit(MutationTypes.SET_IS_FOLLOWING_USER, false)
    } catch (error) {
      throw error
    }
  },
}
