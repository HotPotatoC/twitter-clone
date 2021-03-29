package entity

import "time"

type Tweet struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTweet(id int64, content string, createdAt time.Time) *Tweet {
	return &Tweet{
		ID:        id,
		Content:   content,
		CreatedAt: createdAt,
	}
}
