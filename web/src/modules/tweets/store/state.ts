import { Tweet, TweetAndReplies } from '../types'

export type State = {
  tweetsFeed: Tweet[]
  tweetStatus: TweetAndReplies
  tweetSearchResult: Tweet[]
  tweetImageOverlay: {
    tweetId: number
    show: boolean
    source: string
  }
}

export const state: State = {
  tweetsFeed: [],
  tweetStatus: {
    id: 0,
    content: '',
    photoURLs: null,
    authorName: '',
    authorHandle: '',
    authorPhotoURL: '',
    favoritesCount: 0,
    repliesCount: 0,
    createdAt: '',
    alreadyLiked: false,
    isReply: false,
    replies: [],
  },
  tweetSearchResult: [],
  tweetImageOverlay: {
    tweetId: 0,
    show: false,
    source: '',
  },
}
