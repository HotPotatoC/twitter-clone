import { Module } from 'vuex'
import { state, State } from './state'
import { mutations, Mutations } from './mutations'
import { getters, Getters } from './getters'
import { actions, Actions } from './actions'
import { AugmentedModule } from '../../../store'

export type TweetsModule<S = State> = AugmentedModule<
  S,
  Mutations,
  Getters,
  Actions
>

export const tweetsModule: Module<State, State> = {
  state,
  actions,
  mutations,
  getters,
}
