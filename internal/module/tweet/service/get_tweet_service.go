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
	Name                    string `json:"name"`
	Handle                  string `json:"handle"`
	PhotoURL                string `json:"photo_url"`
	RepliedToTweet          int64  `json:"replied_to_tweet_id,omitempty"`
	RepliedToName           string `json:"replied_to_name,omitempty"`
	RepliedToHandle         string `json:"replied_to_handle,omitempty"`
	RepliedToAuthorPhotoURL string `json:"replied_to_author_photo_url,omitempty"`
	FavoritesCount          int    `json:"favorites_count"`
	RepliesCount            int    `json:"replies_count"`
	AlreadyLiked            bool   `json:"already_liked"`
}

type GetTweetService interface {
	Execute(userID, tweetID int64) (GetTweetOutput, error)
}

type getTweetService struct {
	db database.Database
}

func NewGetTweetService(db database.Database) GetTweetService {
	return getTweetService{db: db}
}

func (s getTweetService) Execute(userID, tweetID int64) (GetTweetOutput, error) {
	var tweetExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM tweets WHERE id = $1)", tweetID).Scan(&tweetExists)
	if err != nil {
		return GetTweetOutput{}, errors.Wrap(err, "service.favoriteTweetService.Execute")
	}

	if !tweetExists {
		return GetTweetOutput{}, entity.ErrTweetDoesNotExist
	}

	var id int64
	var content, name, handle, photoURL string
	var repliedToTweetID sql.NullInt64
	var repliedToName, repliedToHandle, repliedToAuthorPhotoURL sql.NullString
	var createdAt time.Time
	var favoritesCount, repliesCount int
	var alreadyLiked bool

	err = s.db.QueryRow(`
	SELECT
		tweets.id,
		tweets.content,
		tweets.created_at,
		users.name,
		users.handle,
		users.photo_url,
		reply_details.id_tweet,
		reply_details.name,
		reply_details.handle,
		reply_details.photo_url,
		COUNT(favorites.id),
		COUNT(replies.id_reply),
		CASE WHEN favorites.id_user = $1
			AND favorites.id_tweet = tweets.id THEN
			TRUE
		ELSE
			FALSE
		END already_liked
	FROM tweets
		INNER JOIN users ON users.id = tweets.id_user
		LEFT JOIN (
			SELECT replies.id_reply,
				replies.id_tweet,
				users.name,
				users.handle,
				users.photo_url
			FROM replies
				INNER JOIN tweets as t ON t.id = replies.id_tweet
				INNER JOIN users ON users.id = t.id_user
		) as reply_details ON reply_details.id_reply = tweets.id
		LEFT JOIN favorites ON favorites.id_tweet = tweets.id
		LEFT JOIN replies ON replies.id_tweet = tweets.id
	WHERE tweets.id = $2
	GROUP BY
		tweets.id,
		users.name,
		users.handle,
		users.photo_url,
		reply_details.id_tweet,
		reply_details.name,
		reply_details.handle,
		reply_details.photo_url,
		already_liked
	`, userID, tweetID).Scan(&id, &content, &createdAt, &name, &handle, &photoURL, &repliedToTweetID, &repliedToName, &repliedToHandle, &repliedToAuthorPhotoURL, &favoritesCount, &repliesCount, &alreadyLiked)
	if err != nil {
		return GetTweetOutput{}, errors.Wrap(err, "service.getTweetService.Execute")
	}

	return GetTweetOutput{
		Tweet: entity.Tweet{
			ID:        id,
			Content:   content,
			CreatedAt: createdAt,
		},
		Name:                    name,
		Handle:                  handle,
		PhotoURL:                photoURL,
		RepliedToTweet:          repliedToTweetID.Int64,
		RepliedToName:           repliedToName.String,
		RepliedToHandle:         repliedToHandle.String,
		RepliedToAuthorPhotoURL: repliedToAuthorPhotoURL.String,
		FavoritesCount:          favoritesCount,
		RepliesCount:            favoritesCount,
		AlreadyLiked:            alreadyLiked,
	}, nil
}
