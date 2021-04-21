import { ActionTree } from 'vuex'
import { AugmentedActionContext } from '../../../store/types'
import { Mutations, MutationTypes } from './mutations'
import { State, Tweet } from './state'

import axios from '../../../services/axios'

export enum ActionTypes {
  GET_TWEETS_FEED = 'GET_TWEETS_FEED',
  LOAD_MORE_TWEETS = 'LOAD_MORE_TWEETS',
  GET_TWEET_STATUS = 'GET_TWEET_STATUS',
  LOAD_MORE_REPLIES = 'LOAD_MORE_REPLIES',
  SEARCH_TWEETS = 'SEARCH_TWEETS',
  NEW_TWEET = 'NEW_TWEET',
  NEW_REPLY = 'NEW_REPLY',
}

export interface Actions {
  [ActionTypes.GET_TWEETS_FEED]({
    commit,
  }: AugmentedActionContext<Mutations, State>): Promise<any>
  [ActionTypes.LOAD_MORE_TWEETS](
    { commit }: AugmentedActionContext<Mutations, State>,
    cursor: string
  ): Promise<any>
  [ActionTypes.GET_TWEET_STATUS](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: string | string[]
  ): Promise<any>
  [ActionTypes.LOAD_MORE_REPLIES](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { tweetId: string | string[]; cursor: string }
  ): Promise<any>
  [ActionTypes.SEARCH_TWEETS](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: string | string[]
  ): Promise<any>
  [ActionTypes.NEW_TWEET](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: string | string[]
  ): Promise<any>
  [ActionTypes.NEW_REPLY](
    { commit }: AugmentedActionContext<Mutations, State>,
    payload: { tweetId: string | string[]; content: string | string[] }
  ): Promise<any>
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
      const response = await axios.get<TweetsJSONSchema>(`/tweets/feed`)

      const tweetsFeed: Tweet[] = response.data.items.map((item) => ({
        repliedToTweet: item.replied_to_tweet,
        repliedToName: item.replied_to_name,
        favoritesCount: item.favorites_count,
        repliesCount: item.replies_count,
        createdAt: item.created_at,
        alreadyLiked: item.already_liked,
        ...item,
      }))

      commit(MutationTypes.SET_TWEETS_FEED, tweetsFeed)
    } catch (error) {
      return error
    }
  },
  async [ActionTypes.LOAD_MORE_TWEETS]({ commit }, cursor): Promise<any> {
    try {
      const response = await axios.get<TweetsJSONSchema>(
        `/tweets/feed?cursor=${cursor}`
      )

      const tweetsFeed: Tweet[] = response.data.items.map((item) => ({
        repliedToTweet: item.replied_to_tweet,
        repliedToName: item.replied_to_name,
        repliedToHandle: item.replied_to_handle,
        favoritesCount: item.favorites_count,
        repliesCount: item.replies_count,
        createdAt: item.created_at,
        alreadyLiked: item.already_liked,
        ...item,
      }))

      commit(MutationTypes.PUSH_TWEET_FEED, tweetsFeed)
    } catch (error) {
      return error
    }
  },
  async [ActionTypes.GET_TWEET_STATUS]({ commit }, tweetId): Promise<any> {
    try {
      const tweetResponse = await axios.get<TweetJSONSchema>(
        `/tweets/${tweetId}`
      )
      const repliesResponse = await axios.get<TweetsJSONSchema>(
        `/tweets/${tweetId}/replies`
      )

      commit(MutationTypes.SET_TWEET_STATUS, {
        repliedToTweet: tweetResponse.data.replied_to_tweet,
        repliedToName: tweetResponse.data.replied_to_name,
        favoritesCount: tweetResponse.data.favorites_count,
        repliesCount: tweetResponse.data.replies_count,
        createdAt: tweetResponse.data.created_at,
        replies:
          repliesResponse.data.items !== null
            ? repliesResponse.data.items.map((item) => ({
                repliedToTweet: item.replied_to_tweet,
                repliedToName: item.created_at,
                favoritesCount: item.favorites_count,
                repliesCount: item.replies_count,
                createdAt: item.created_at,
                ...item,
              }))
            : [],
        ...tweetResponse.data,
      })
    } catch (error) {
      return error
    }
  },
  async [ActionTypes.LOAD_MORE_REPLIES](
    { commit },
    { tweetId, cursor }
  ): Promise<any> {
    try {
      const response = await axios.get<TweetsJSONSchema>(
        `/tweets/${tweetId}/replies${cursor ? `?cursor=${cursor}` : ''}`
      )

      if (response.data.items === null) {
        return
      }

      const replies: Tweet[] = response.data.items.map((item) => ({
        favoritesCount: item.favorites_count,
        repliesCount: item.replies_count,
        createdAt: item.created_at,
        ...item,
      }))

      commit(MutationTypes.PUSH_REPLIES_TO_TWEET_STATUS, replies)
    } catch (error) {
      return error
    }
  },
  async [ActionTypes.SEARCH_TWEETS]({ commit }, query): Promise<any> {
    try {
      const response = await axios.get<TweetsJSONSchema>(
        `/tweets/search?query=${query}`
      )

      const searchResults: Tweet[] = response.data.items.map((item) => ({
        repliedToTweet: item.replied_to_tweet,
        repliedToName: item.replied_to_name,
        repliedToHandle: item.replied_to_handle,
        favoritesCount: item.favorites_count,
        repliesCount: item.replies_count,
        createdAt: item.created_at,
        ...item,
      }))

      commit(MutationTypes.SET_TWEET_SEARCH_RESULTS, searchResults)
    } catch (error) {
      return error
    }
  },
  async [ActionTypes.NEW_TWEET]({ commit }, content): Promise<any> {
    try {
      await axios.post<TweetsJSONSchema>('/tweets', { content })
    } catch (error) {
      return error
    }
  },
  async [ActionTypes.NEW_REPLY](
    { commit },
    { tweetId, content }
  ): Promise<any> {
    try {
      await axios.post<TweetsJSONSchema>(`/tweets/${tweetId}/reply`, {
        content,
      })
    } catch (error) {
      return error
    }
  },
}
