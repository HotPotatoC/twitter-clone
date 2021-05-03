export type TweetJSONSchema = {
  id: number
  content: string
  photo_urls: string[] | null
  author_name: string
  author_handle: string
  author_photo_url: string
  replied_to?: Omit<TweetJSONSchema, 'replied_to' | 'is_reply' | 'created_at'>
  is_reply: boolean
  favorites_count: number
  replies_count: number
  created_at: string
  already_liked: boolean
}

export type Tweet = {
  id: number
  content: string
  photoURLs: string[] | null
  authorName: string
  authorHandle: string
  authorPhotoURL: string
  repliedTo?: Omit<Tweet, 'repliedTo' | 'isReply' | 'createdAt'>
  favoritesCount: number
  repliesCount: number
  isReply: boolean
  alreadyLiked?: boolean
  createdAt: string
}

export type TweetAndReplies = Tweet & {
  replies: Tweet[]
}
