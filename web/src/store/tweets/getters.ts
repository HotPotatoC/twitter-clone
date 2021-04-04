import { GetterTree } from 'vuex'
import { State, TweetsFeed } from './state'

export interface Getters {
  getTweetsFeed(state: State): TweetsFeed[]
}

export const getters: GetterTree<State, State> & Getters = {
  getTweetsFeed(state): TweetsFeed[] {
    return state.tweetsFeed
  },
}
