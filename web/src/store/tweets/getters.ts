import { GetterTree } from 'vuex'
import { State, Tweet } from './state'

export interface Getters {
  getTweetsFeed(state: State): Tweet[]
  getTweetStatus(state: State): Tweet & { replies: Tweet[] }
}

export const getters: GetterTree<State, State> & Getters = {
  getTweetsFeed(state): Tweet[] {
    return state.tweetsFeed
  },
  getTweetStatus(state): Tweet & { replies: Tweet[] } {
    return state.tweetStatus
  },
}
