package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type GetUserTweetsOutput struct {
	entity.Tweet
	Name            string `json:"name"`
	Handle          string `json:"handle"`
	RepliedToTweet  int64  `json:"replied_to_tweet_id,omitempty"`
	RepliedToName   string `json:"replied_to_name,omitempty"`
	RepliedToHandle string `json:"replied_to_handle,omitempty"`
	FavoritesCount  int    `json:"favorites_count"`
	RepliesCount    int    `json:"replies_count"`
}

type GetUserTweetsService interface {
	Execute(username string, createdAtCursor string) ([]GetUserTweetsOutput, error)
}

type getUserTweetsService struct {
	db database.Database
}

func NewGetUserTweetsService(db database.Database) GetUserTweetsService {
	return getUserTweetsService{db: db}
}

func (s getUserTweetsService) Execute(username string, createdAtCursor string) ([]GetUserTweetsOutput, error) {
	var tweets []GetUserTweetsOutput

	var rows database.Rows
	var err error

	withCursor := createdAtCursor != ""
	query := s.buildSQLQuery(withCursor)

	if withCursor {
		cursor, err := time.Parse(time.RFC3339, createdAtCursor)
		if err != nil {
			return []GetUserTweetsOutput{}, ErrInvalidCursor
		}

		rows, err = s.db.Query(query, username, cursor)
		if err != nil {
			return []GetUserTweetsOutput{}, errors.Wrap(err, "service.getUserTweetsService.Execute")
		}
	} else {
		rows, err = s.db.Query(query, username)
		if err != nil {
			return []GetUserTweetsOutput{}, errors.Wrap(err, "service.getUserTweetsService.Execute")
		}
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var content, name, handle string
		var repliedToTweetID sql.NullInt64
		var repliedToName, repliedToHandle sql.NullString
		var createdAt time.Time
		var favoritesCount, repliesCount int

		err = rows.Scan(&id, &content, &createdAt, &name, &handle, &repliedToTweetID, &repliedToName, &repliedToHandle, &favoritesCount, &repliesCount)
		if err != nil {
			return []GetUserTweetsOutput{}, errors.Wrap(err, "service.getUserTweetsService.Execute")
		}

		tweets = append(tweets, GetUserTweetsOutput{
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
			RepliesCount:    repliesCount,
		})
	}

	if err := rows.Err(); err != nil {
		return []GetUserTweetsOutput{}, errors.Wrap(err, "service.getUserTweetsService.Execute")
	}

	return tweets, nil
}

func (s getUserTweetsService) buildSQLQuery(withCursor bool) string {
	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
	SELECT
		tweets.id,
		tweets.content,
		tweets.created_at,
		(ARRAY_AGG(users.name))[1],
		(ARRAY_AGG(users.handle))[1],
		(ARRAY_AGG(sq.id_tweet))[1],
		(ARRAY_AGG(sq.name))[1],
		(ARRAY_AGG(sq.handle))[1],
		COUNT(f.*),
		COUNT(r1.*)
	FROM
		tweets
		LEFT JOIN users ON users.id = tweets.id_user
		LEFT JOIN (
			SELECT
				replies.id_reply,
				replies.id_tweet,
				users.name,
				users.handle
			FROM
				replies
				INNER JOIN tweets AS t ON t.id = replies.id_tweet
				INNER JOIN users ON users.id = t.id_user) AS sq ON sq.id_reply = tweets.id
		LEFT JOIN favorites AS f ON f.id_tweet = tweets.id
		LEFT JOIN replies AS r1 ON r1.id_tweet = tweets.id
	WHERE users.handle = $1
	`)

	if withCursor {
		queryBuilder.WriteString("AND tweets.created_at < $2")
	}

	queryBuilder.WriteString(`
	GROUP BY
		tweets.id
	ORDER BY
		tweets.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
