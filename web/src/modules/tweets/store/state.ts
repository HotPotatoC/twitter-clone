export interface Tweet {
  id: number
  content: string
  name: string
  handle: string
  repliedToTweet?: number
  repliedToName?: string
  repliedToHandle?: string
  favoritesCount: number
  repliesCount: number
  alreadyLiked?: boolean
  createdAt: string
}

export interface TweetAndReplies extends Tweet {
  replies: Tweet[]
}

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
