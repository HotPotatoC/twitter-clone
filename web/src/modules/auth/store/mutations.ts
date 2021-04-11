import { MutationTree } from 'vuex'
import { AuthStatus, State, UserData } from './state'

export enum MutationTypes {
  SET_AUTHENTICATION_STATUS = 'SET_AUTHENTICATION_STATUS',
  SET_ACCESS_TOKEN = 'SET_ACCESS_TOKEN',
  SET_USER_DATA = 'SET_USER_DATA',
}

export interface Mutations<S = State> {
  [MutationTypes.SET_AUTHENTICATION_STATUS](state: S, payload: AuthStatus): void
  [MutationTypes.SET_ACCESS_TOKEN](state: S, payload: string): void
  [MutationTypes.SET_USER_DATA](state: S, payload: UserData): void
}

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_AUTHENTICATION_STATUS](state, payload) {
    state.authStatus = payload
  },
  [MutationTypes.SET_ACCESS_TOKEN](state, payload) {
    state.accessToken = payload
  },
  [MutationTypes.SET_USER_DATA](state, payload) {
    state.user = payload
  },
}
