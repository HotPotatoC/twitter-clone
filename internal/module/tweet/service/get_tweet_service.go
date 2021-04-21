package service

import (
	"database/sql"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type GetTweetOutput struct {
	entity.Tweet
	Name            string `json:"name"`
	Handle          string `json:"handle"`
	RepliedToTweet  int64  `json:"replied_to_tweet_id,omitempty"`
	RepliedToName   string `json:"replied_to_name,omitempty"`
	RepliedToHandle string `json:"replied_to_handle,omitempty"`
	FavoritesCount  int    `json:"favorites_count"`
	RepliesCount    int    `json:"replies_count"`
}

type GetTweetService interface {
	Execute(tweetID int64) (GetTweetOutput, error)
}

type getTweetService struct {
	db database.Database
}

func NewGetTweetService(db database.Database) GetTweetService {
	return getTweetService{db: db}
}

func (s getTweetService) Execute(tweetID int64) (GetTweetOutput, error) {
	var tweetExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM tweets WHERE id = $1)", tweetID).Scan(&tweetExists)
	if err != nil {
		return GetTweetOutput{}, errors.Wrap(err, "service.favoriteTweetService.Execute")
	}

	if !tweetExists {
		return GetTweetOutput{}, entity.ErrTweetDoesNotExist
	}

	var id int64
	var content, name, handle string
	var repliedToTweetID sql.NullInt64
	var repliedToName, repliedToHandle sql.NullString
	var createdAt time.Time
	var favoritesCount, repliesCount int

	err = s.db.QueryRow(`
	SELECT tweets.id,
		tweets.content,
		tweets.created_at,
		users.name,
		users.handle,
		reply_details.id_tweet,
		reply_details.name,
		reply_details.handle,
		COUNT(favorites.id),
		COUNT(replies.id_reply)
	FROM tweets
		LEFT JOIN users ON users.id = tweets.id_user
		LEFT JOIN (
			SELECT replies.id_reply,
				replies.id_tweet,
				users.name,
				users.handle
			FROM replies
				INNER JOIN tweets as t ON t.id = replies.id_tweet
				INNER JOIN users ON users.id = t.id_user
		) as reply_details ON reply_details.id_reply = tweets.id
		LEFT JOIN favorites ON favorites.id_tweet = tweets.id
		LEFT JOIN replies ON replies.id_tweet = tweets.id
	WHERE tweets.id = $1
	GROUP BY
		tweets.id,
		users.name,
		users.handle,
		reply_details.id_tweet,
		reply_details.name,
		reply_details.handle
	`, tweetID).Scan(&id, &content, &createdAt, &name, &handle, &repliedToTweetID, &repliedToName, &repliedToHandle, &favoritesCount, &repliesCount)
	if err != nil {
		return GetTweetOutput{}, errors.Wrap(err, "service.getTweetService.Execute")
	}

	return GetTweetOutput{
		Tweet: entity.Tweet{
			ID:        id,
			Content:   content,
			CreatedAt: createdAt,
		},
		Name:            name,
		Handle:          handle,
		RepliedToTweet:  repliedToTweetID.Int64,
		RepliedToName:   repliedToName.String,
		RepliedToHandle: repliedToHandle.String,
		FavoritesCount:  favoritesCount,
		RepliesCount:    favoritesCount,
	}, nil
}
