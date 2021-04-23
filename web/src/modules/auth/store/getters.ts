import { GetterTree } from 'vuex'
import { State } from './state'
import { UserData } from '../types'

export type Getters = {
  isLoggedIn(state: State): boolean
  accessToken(state: State): string
  userData(state: State): UserData
}

export const getters: GetterTree<State, State> & Getters = {
  isLoggedIn(state): boolean {
    return state.authStatus.isLoggedIn
  },
  accessToken(state): string {
    return state.accessToken
  },
  userData(state): UserData {
    return state.user
  },
}
