export type TweetJSONSchema = {
  id: number
  content: string
  name: string
  handle: string
  photo_url: string
  replied_to_tweet: number
  replied_to_name: string
  replied_to_handle: string
  replied_to_photo_url: string
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
  photoURL: string
  repliedToTweet?: number
  repliedToName?: string
  repliedToHandle?: string
  repliedToPhotoURL?: string
  favoritesCount: number
  repliesCount: number
  alreadyLiked?: boolean
  createdAt: string
}

export type TweetAndReplies = Tweet & {
  replies: Tweet[]
}
