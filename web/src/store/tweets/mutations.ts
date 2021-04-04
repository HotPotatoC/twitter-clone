import { MutationTree } from 'vuex'
import { State, TweetsFeed } from './state'

export enum MutationTypes {
  SET_TWEETS_FEED = 'SET_TWEETS_FEED',
}

export interface Mutations<S = State> {
  [MutationTypes.SET_TWEETS_FEED](state: S, payload: TweetsFeed[]): void
}

export const mutations: MutationTree<State> & Mutations = {
  [MutationTypes.SET_TWEETS_FEED](state, payload) {
    state.tweetsFeed = payload
  },
}
