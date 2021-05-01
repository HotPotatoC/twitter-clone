import { Tweet, TweetAndReplies } from '../types'

export type State = {
  tweetsFeed: Tweet[]
  tweetStatus: TweetAndReplies
  tweetSearchResult: Tweet[]
}

export const state: State = {
  tweetsFeed: [],
  tweetStatus: {
    id: 0,
    content: '',
    authorName: '',
    authorHandle: '',
    authorPhotoURL: '',
    favoritesCount: 0,
    repliesCount: 0,
    createdAt: '',
    alreadyLiked: false,
    replies: [],
  },
  tweetSearchResult: [],
}
