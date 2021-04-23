import { MutationTree } from 'vuex'
import { State } from './state'
import { Theme } from '../types'

export enum MutationTypes {
  SET_THEME = 'SET_THEME',
}

export type Mutations<S = State> = {
  [MutationTypes.SET_THEME](state: S, payload: Theme): void
}

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_THEME](state, payload) {
    state.theme = payload
  },
}
