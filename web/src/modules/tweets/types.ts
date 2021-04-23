export type TweetJSONSchema = {
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
  replies?: TweetJSONSchema[]
}

export type Tweet = {
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

export type TweetAndReplies = Tweet & {
  replies: Tweet[]
}
