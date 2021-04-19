import { MutationTree } from 'vuex'
import { State, ProfileDetails, ProfileStatus } from './state'

export enum MutationTypes {
  SET_PROFILE_DETAILS = 'SET_PROFILE_DETAILS',
  SET_PROFILE_STATUS = 'SET_PROFILE_STATUS',
}

export interface Mutations<S = State> {
  [MutationTypes.SET_PROFILE_DETAILS](state: S, payload: ProfileDetails): void
  [MutationTypes.SET_PROFILE_STATUS](state: S, payload: ProfileStatus): void
}

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_PROFILE_DETAILS](state, payload) {
    state.profileDetails = payload
  },
  [MutationTypes.SET_PROFILE_STATUS](state, payload) {
    state.status = payload
  },
}
