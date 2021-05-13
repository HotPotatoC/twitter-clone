package entity

import "time"

type Tweet struct {
	ID             int64     `json:"id"`
	Content        string    `json:"content"`
	PhotoURLs      []string  `json:"photo_urls"`
	FavoritesCount int       `json:"favorites_count"`
	RepliesCount   int       `json:"replies_count"`
	RetweetsCount  int       `json:"retweets_count"`
	CreatedAt      time.Time `json:"created_at"`
}

type Reply struct {
	ID             int64    `json:"id,omitempty"`
	Content        string   `json:"content,omitempty"`
	PhotoURLs      []string `json:"photo_urls"`
	AuthorName     string   `json:"author_name,omitempty"`
	AuthorHandle   string   `json:"author_handle,omitempty"`
	AuthorPhotoURL string   `json:"author_photo_url,omitempty"`
	FavoritesCount int      `json:"favorites_count"`
	RepliesCount   int      `json:"replies_count"`
	RetweetsCount  int      `json:"retweets_count"`
	AlreadyLiked   bool     `json:"already_liked"`
}
