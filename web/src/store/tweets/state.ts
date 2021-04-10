export interface Tweet {
  id: number
  content: string
  name: string
  repliedToTweet?: number
  repliedToName?: string
  favoritesCount: number
  repliesCount: number
  createdAt: string
}

export interface TweetAndReplies extends Tweet {
  replies: Tweet[]
}

export interface State {
  tweetsFeed: Tweet[]
  tweetStatus: TweetAndReplies
}

export const state: State = {
  tweetsFeed: [],
  tweetStatus: {
    id: 0,
    content: '',
    name: '',
    repliedToTweet: 0,
    repliedToName: '',
    favoritesCount: 0,
    repliesCount: 0,
    createdAt: '',
    replies: [],
  },
}
