import { ActionTree } from 'vuex'
import { AugmentedActionContext } from '../../../store'
import { Mutations, MutationTypes } from './mutations'
import { State } from './state'

import {
  createReply,
  createTweet,
  favoriteTweet,
  fetchFeed,
  fetchReplies,
  fetchTweet,
  searchTweets,
  retweet,
} from '../service'
import { Tweet } from '../types'

export enum ActionTypes {
  GET_TWEETS_FEED = 'GET_TWEETS_FEED',
  LOAD_MORE_TWEETS = 'LOAD_MORE_TWEETS',
  GET_TWEET_STATUS = 'GET_TWEET_STATUS',
  LOAD_MORE_REPLIES = 'LOAD_MORE_REPLIES',
  SEARCH_TWEETS = 'SEARCH_TWEETS',
  NEW_TWEET = 'NEW_TWEET',
  NEW_REPLY = 'NEW_REPLY',
  FAVORITE_TWEET = 'FAVORITE_TWEET',
  TOGGLE_TWEET_IMAGE_OVERLAY = 'TOGGLE_TWEET_IMAGE_OVERLAY',
  RETWEET = 'RETWEET',
}

export type Actions = {
  [ActionTypes.GET_TWEETS_FEED]({
    commit,
  }: AugmentedActionContext<Mutations, State>): Promise<any>
  [ActionTypes.LOAD_MORE_TWEETS](
    { commit }: AugmentedActionContext<Mutations, State>,
    cursor: string
  ): Promise<any>
  [ActionTypes.GET_TWEET_STATUS](
    { commit }: AugmentedActionContext<Mutations, State>,
    tweetId: string
  ): Promise<any>
  [ActionTypes.LOAD_MORE_REPLIES](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { tweetId: string; cursor: string }
  ): Promise<any>
  [ActionTypes.SEARCH_TWEETS](
    { commit }: AugmentedActionContext<Mutations, State>,
    query: string
  ): Promise<any>
  [ActionTypes.NEW_TWEET](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { content: string; attachments?: File[] }
  ): Promise<any>
  [ActionTypes.NEW_REPLY](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { tweetId: string; content: string }
  ): Promise<any>
  [ActionTypes.FAVORITE_TWEET](
    { commit }: AugmentedActionContext<Mutations, State>,
    tweetId: string
  ): Promise<any>
  [ActionTypes.TOGGLE_TWEET_IMAGE_OVERLAY](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { tweet: Tweet; show: boolean; source: string }
  ): any
  [ActionTypes.RETWEET](
    { commit }: AugmentedActionContext<Mutations, State>,
    tweetId: string
  ): any
}

export interface TweetJSONSchema {
  id: number
  content: string
  name: string
  handle: string
  replied_to_tweet: number
  replied_to_name: string
  replied_to_handle: string
  favorites_count: number
  replies_count: number
  created_at: string
  already_liked: boolean
}

export interface TweetsJSONSchema {
  items: TweetJSONSchema[]
  total_records: number
}

export const actions: ActionTree<State, State> & Actions = {
  async [ActionTypes.GET_TWEETS_FEED]({ commit }): Promise<any> {
    try {
      const feed = await fetchFeed()

      commit(MutationTypes.SET_TWEETS_FEED, feed)
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.LOAD_MORE_TWEETS]({ commit }, cursor): Promise<any> {
    try {
      const feed = await fetchFeed(cursor)

      commit(MutationTypes.PUSH_TWEET_FEED, feed)
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.GET_TWEET_STATUS]({ commit }, tweetId): Promise<any> {
    try {
      const tweet = await fetchTweet(tweetId)
      const replies = await fetchReplies(tweetId)

      commit(MutationTypes.SET_TWEET_STATUS, { ...tweet, replies })
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.LOAD_MORE_REPLIES](
    { commit },
    { tweetId, cursor }
  ): Promise<any> {
    try {
      const replies = await fetchReplies(tweetId, cursor)

      if (replies === null) {
        return
      }

      commit(MutationTypes.PUSH_REPLIES_TO_TWEET_STATUS, replies)
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.SEARCH_TWEETS]({ commit }, query): Promise<any> {
    try {
      const results = await searchTweets(query)

      commit(MutationTypes.SET_TWEET_SEARCH_RESULTS, results)
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.NEW_TWEET](
    { commit },
    { content, attachments }
  ): Promise<any> {
    try {
      await createTweet(content, attachments)
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.NEW_REPLY](
    { commit },
    { tweetId, content }
  ): Promise<any> {
    try {
      await createReply(tweetId, content)
    } catch (error) {
      throw error
    }
  },
  async [ActionTypes.FAVORITE_TWEET]({ commit }, tweetId): Promise<any> {
    try {
      await favoriteTweet(tweetId)
    } catch (error) {
      throw error
    }
  },
  [ActionTypes.TOGGLE_TWEET_IMAGE_OVERLAY](
    { commit },
    { tweet, show, source }
  ) {
    commit(MutationTypes.TOGGLE_TWEET_IMAGE_OVERLAY, {
      tweet,
      show,
      source,
    })
  },
  async [ActionTypes.RETWEET]({ commit }, tweetId): Promise<any> {
    try {
      await retweet(tweetId)
    } catch (error) {
      throw error
    }
  },
}
