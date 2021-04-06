import { GetterTree } from 'vuex'
import { State, Tweet, TweetAndReplies } from './state'

export interface Getters {
  getTweetsFeed(state: State): Tweet[]
  getTweetStatus(state: State): TweetAndReplies
}

export const getters: GetterTree<State, State> & Getters = {
  getTweetsFeed(state): Tweet[] {
    return state.tweetsFeed
  },
  getTweetStatus(state): TweetAndReplies {
    return state.tweetStatus
  },
}
