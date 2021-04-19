import { GetterTree } from 'vuex'
import { State, ProfileDetails } from './state'

export interface Getters {
  profileInfo(state: State): ProfileDetails
}

export const getters: GetterTree<State, State> & Getters = {
  profileInfo(state): ProfileDetails {
    return state.profileDetails
  },
}
