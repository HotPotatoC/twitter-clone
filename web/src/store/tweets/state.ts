export interface TweetsFeed {
  id: number
  content: string
  name: string
  repliedToTweet: number
  repliedToName: string
  favoritesCount: number
  repliesCount: number
  createdAt: string
}

export interface State {
  tweetsFeed: TweetsFeed[]
}

export const state: State = {
  tweetsFeed: [],
}
