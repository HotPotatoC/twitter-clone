import { Tweet, TweetAndReplies } from '../types'

export interface State {
  tweetsFeed: Tweet[]
  tweetStatus: TweetAndReplies
  tweetSearchResult: Tweet[]
}

export const state: State = {
  tweetsFeed: [],
  tweetStatus: {
    id: 0,
    content: '',
    name: '',
    handle: '',
    repliedToTweet: 0,
    repliedToName: '',
    repliedToHandle: '',
    favoritesCount: 0,
    repliesCount: 0,
    createdAt: '',
    alreadyLiked: false,
    replies: [],
  },
  tweetSearchResult: [],
}
