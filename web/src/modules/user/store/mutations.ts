import { MutationTree } from 'vuex'
import { Tweet } from '../../tweets/types'
import { State } from './state'
import { ProfileStatus, ProfileDetails } from '../types'

export enum MutationTypes {
  SET_PROFILE_STATUS = 'SET_PROFILE_STATUS',
  SET_PROFILE_DETAILS = 'SET_PROFILE_DETAILS',
  SET_PROFILE_TWEETS = 'SET_PROFILE_TWEETS',
  PUSH_PROFILE_TWEETS = 'PUSH_PROFILE_TWEETS',
}

export type Mutations<S = State> = {
  [MutationTypes.SET_PROFILE_STATUS](state: S, payload: ProfileStatus): void
  [MutationTypes.SET_PROFILE_DETAILS](state: S, payload: ProfileDetails): void
  [MutationTypes.SET_PROFILE_TWEETS](state: S, payload: Tweet[]): void
  [MutationTypes.PUSH_PROFILE_TWEETS](state: S, payload: Tweet[]): void
}

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_PROFILE_STATUS](state, payload) {
    state.status = payload
  },
  [MutationTypes.SET_PROFILE_DETAILS](state, payload) {
    state.profileDetails = payload
  },
  [MutationTypes.SET_PROFILE_TWEETS](state, payload) {
    state.profileTweets = payload
  },
  [MutationTypes.PUSH_PROFILE_TWEETS](state, payload) {
    state.profileTweets.push(...payload)
  },
}
