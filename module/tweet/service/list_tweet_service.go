package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/database"
	"github.com/HotPotatoC/twitter-clone/module"
	"github.com/HotPotatoC/twitter-clone/module/tweet/entity"
	"github.com/pkg/errors"
)

type ListTweetOutput struct {
	entity.Tweet
	Name              string `json:"name"`
	Handle            string `json:"handle"`
	PhotoURL          string `json:"photo_url"`
	RepliedToTweet    int64  `json:"replied_to_tweet_id,omitempty"`
	RepliedToName     string `json:"replied_to_name,omitempty"`
	RepliedToHandle   string `json:"replied_to_handle,omitempty"`
	RepliedToPhotoURL string `json:"replied_to_photo_url,omitempty"`
	FavoritesCount    int    `json:"favorites_count"`
	RepliesCount      int    `json:"replies_count"`
}

type ListTweetService interface {
	Execute(createdAtCursor string) ([]ListTweetOutput, error)
}

type listTweetService struct {
	db database.Database
}

func NewListTweetService(db database.Database) ListTweetService {
	return listTweetService{db: db}
}

func (s listTweetService) Execute(createdAtCursor string) ([]ListTweetOutput, error) {
	var tweets []ListTweetOutput

	var rows database.Rows
	var err error

	withCursor := createdAtCursor != ""
	query := s.buildSQLQuery(withCursor)

	if withCursor {
		cursor, err := time.Parse(time.RFC3339, createdAtCursor)
		if err != nil {
			return []ListTweetOutput{}, module.ErrInvalidCursor
		}

		rows, err = s.db.Query(query, cursor)
		if err != nil {
			return []ListTweetOutput{}, errors.Wrap(err, "service.listTweetService.Execute")
		}
	} else {
		rows, err = s.db.Query(query)
		if err != nil {
			return []ListTweetOutput{}, errors.Wrap(err, "service.listTweetService.Execute")
		}
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var content, name, handle, photoURL string
		var repliedToTweetID sql.NullInt64
		var repliedToName, repliedToHandle, repliedToPhotoURL sql.NullString
		var createdAt time.Time
		var favoritesCount, repliesCount int

		err = rows.Scan(&id, &content, &createdAt, &name, &handle, &photoURL, &repliedToTweetID, &repliedToName, &repliedToHandle, &repliedToPhotoURL, &favoritesCount, &repliesCount)
		if err != nil {
			return []ListTweetOutput{}, errors.Wrap(err, "service.listTweetService.Execute")
		}

		tweets = append(tweets, ListTweetOutput{
			Tweet: entity.Tweet{
				ID:        id,
				Content:   content,
				CreatedAt: createdAt,
			},
			Name:              name,
			Handle:            handle,
			PhotoURL:          photoURL,
			RepliedToTweet:    repliedToTweetID.Int64,
			RepliedToName:     repliedToName.String,
			RepliedToHandle:   repliedToHandle.String,
			RepliedToPhotoURL: repliedToPhotoURL.String,
			FavoritesCount:    favoritesCount,
			RepliesCount:      repliesCount,
		})
	}

	if err := rows.Err(); err != nil {
		return []ListTweetOutput{}, errors.Wrap(err, "service.listTweetService.Execute")
	}

	return tweets, nil
}

func (s listTweetService) buildSQLQuery(withCursor bool) string {
	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
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
		COUNT(replies.id_reply)
	FROM tweets
		LEFT JOIN users ON users.id = tweets.id_user
		LEFT JOIN (
			SELECT
				replies.id_reply,
				replies.id_tweet,
				users.name,
				users.handle,
				users.photo_url
			FROM replies
				INNER JOIN tweets AS t ON t.id = replies.id_tweet
				INNER JOIN users ON users.id = t.id_user) AS reply_details ON reply_details.id_reply = tweets.id
		LEFT JOIN favorites ON favorites.id_tweet = tweets.id
		LEFT JOIN replies ON replies.id_tweet = tweets.id
	`)

	if withCursor {
		queryBuilder.WriteString("WHERE tweets.created_at < $1")
	}

	queryBuilder.WriteString(`
	GROUP BY
		tweets.id,
		users.name,
		users.handle,
		users.photo_url,
		reply_details.id_tweet,
		reply_details.name,
		reply_details.handle,
		reply_details.photo_url
	ORDER BY
		tweets.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
