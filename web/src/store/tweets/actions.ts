import { ActionTree } from 'vuex'
import { AugmentedActionContext } from '../types'
import { Mutations, MutationTypes } from './mutations'
import { State, TweetsFeed } from './state'

import axios from '../../services/axios'

export enum ActionTypes {
  GET_TWEETS_FEED = 'GET_TWEETS_FEED',
}

export interface Actions {
  [ActionTypes.GET_TWEETS_FEED]({
    commit,
  }: AugmentedActionContext<Mutations, State>): Promise<any>
}

interface GetTweetsFeedResponse {
  items: {
    id: number
    content: string
    name: string
    replied_to_tweet: number
    replied_to_name: string
    favorites_count: number
    replies_count: number
    created_at: string
  }[]
  total_records: number
}

export const actions: ActionTree<State, State> & Actions = {
  async [ActionTypes.GET_TWEETS_FEED]({ commit }): Promise<any> {
    try {
      const response = await axios.get<GetTweetsFeedResponse>('/tweets')

      const tweetsFeed: TweetsFeed[] = response.data.items.map((item) => ({
        repliedToTweet: item.replied_to_tweet,
        repliedToName: item.created_at,
        favoritesCount: item.favorites_count,
        repliesCount: item.replies_count,
        createdAt: item.created_at,
        ...item,
      }))

      commit(MutationTypes.SET_TWEETS_FEED, tweetsFeed)
    } catch (error) {
      return error
    }
  },
}
