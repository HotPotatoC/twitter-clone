import { GetterTree } from 'vuex'
import { State, Tweet, TweetAndReplies } from './state'

export interface Getters {
  getTweetsFeed(state: State): Tweet[]
  getLastTweetFeedItem(state: State): Tweet
  getTweetStatus(state: State): TweetAndReplies
}

export const getters: GetterTree<State, State> & Getters = {
  getTweetsFeed(state): Tweet[] {
    return state.tweetsFeed
  },
  getLastTweetFeedItem(state): Tweet {
    return state.tweetsFeed[state.tweetsFeed.length - 1]
  },
  getTweetStatus(state): TweetAndReplies {
    return state.tweetStatus
  },
}
