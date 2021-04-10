import { MutationTree } from 'vuex'
import { State, Tweet, TweetAndReplies } from './state'

export enum MutationTypes {
  SET_TWEETS_FEED = 'SET_TWEETS_FEED',
  PUSH_TWEET_FEED = 'PUSH_TWEET_FEED',
  SET_TWEET_STATUS = 'SET_TWEET_STATUS',
}

export interface Mutations<S = State> {
  [MutationTypes.SET_TWEETS_FEED](state: S, payload: Tweet[]): void
  [MutationTypes.PUSH_TWEET_FEED](state: S, payload: Tweet[]): void
  [MutationTypes.SET_TWEET_STATUS](state: S, payload: TweetAndReplies): void
}

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_TWEETS_FEED](state, payload) {
    state.tweetsFeed = payload
  },
  [MutationTypes.PUSH_TWEET_FEED](state, payload) {
    state.tweetsFeed.push(...payload)
  },
  [MutationTypes.SET_TWEET_STATUS](state, payload) {
    state.tweetStatus = payload
  },
}
