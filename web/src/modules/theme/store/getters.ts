import { GetterTree } from 'vuex'
import { State } from './state'

export type Getters = {
  currentTheme(state: State): string
}

export const getters: GetterTree<State, State> & Getters = {
  currentTheme(state): string {
    return state.theme
  },
}
