import { MutationTree } from 'vuex'
import { State } from './state'
import { Tweet, TweetAndReplies } from '../types'

export enum MutationTypes {
  SET_TWEETS_FEED = 'SET_TWEETS_FEED',
  PUSH_TWEET_FEED = 'PUSH_TWEET_FEED',
  SET_TWEET_STATUS = 'SET_TWEET_STATUS',
  PUSH_REPLIES_TO_TWEET_STATUS = 'PUSH_REPLIES_TO_TWEET_STATUS',
  SET_TWEET_SEARCH_RESULTS = 'SET_TWEET_SEARCH_RESULTS',
  PUSH_TWEET_SEARCH_RESULTS = 'PUSH_TWEET_SEARCH_RESULTS',
}

export type Mutations<S = State> = {
  [MutationTypes.SET_TWEETS_FEED](state: S, payload: Tweet[]): void
  [MutationTypes.PUSH_TWEET_FEED](state: S, payload: Tweet[]): void
  [MutationTypes.SET_TWEET_STATUS](state: S, payload: TweetAndReplies): void
  [MutationTypes.PUSH_REPLIES_TO_TWEET_STATUS](state: S, payload: Tweet[]): void
  [MutationTypes.SET_TWEET_SEARCH_RESULTS](state: S, payload: Tweet[]): void
  [MutationTypes.PUSH_TWEET_SEARCH_RESULTS](state: S, payload: Tweet[]): void
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
  [MutationTypes.PUSH_REPLIES_TO_TWEET_STATUS](state, payload) {
    state.tweetStatus.replies.push(...payload)
  },
  [MutationTypes.SET_TWEET_SEARCH_RESULTS](state, payload) {
    state.tweetSearchResult = payload
  },
  [MutationTypes.PUSH_TWEET_SEARCH_RESULTS](state, payload) {
    state.tweetSearchResult.push(...payload)
  },
}
