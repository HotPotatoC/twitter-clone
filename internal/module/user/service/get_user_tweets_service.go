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
	AlreadyLiked    bool   `json:"already_liked"`
}

type GetUserTweetsService interface {
	Execute(userID int64, username string, createdAtCursor string) ([]GetUserTweetsOutput, error)
}

type getUserTweetsService struct {
	db database.Database
}

func NewGetUserTweetsService(db database.Database) GetUserTweetsService {
	return getUserTweetsService{db: db}
}

func (s getUserTweetsService) Execute(userID int64, username string, createdAtCursor string) ([]GetUserTweetsOutput, error) {
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

		rows, err = s.db.Query(query, userID, username, cursor)
		if err != nil {
			return []GetUserTweetsOutput{}, errors.Wrap(err, "service.getUserTweetsService.Execute")
		}
	} else {
		rows, err = s.db.Query(query, userID, username)
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
		var alreadyLiked bool

		err = rows.Scan(&id, &content, &createdAt, &name, &handle, &repliedToTweetID, &repliedToName, &repliedToHandle, &favoritesCount, &repliesCount, &alreadyLiked)
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
			AlreadyLiked:    alreadyLiked,
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
		users.name,
		users.handle,
		reply_details.id_tweet,
		reply_details.name,
		reply_details.handle,
		COUNT(favorites.id),
		COUNT(replies.id_reply),
		CASE WHEN favorites.id_user = $1
			AND favorites.id_tweet = tweets.id THEN
			TRUE
		ELSE
			FALSE
		END already_liked
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
				INNER JOIN users ON users.id = t.id_user) AS reply_details ON reply_details.id_reply = tweets.id
		LEFT JOIN favorites ON favorites.id_tweet = tweets.id
		LEFT JOIN replies ON replies.id_tweet = tweets.id
	WHERE users.handle = $2
	`)

	if withCursor {
		queryBuilder.WriteString("AND tweets.created_at < $3")
	}

	queryBuilder.WriteString(`
	GROUP BY
		tweets.id,
		users.name,
		users.handle,
		reply_details.id_tweet,
		reply_details.name,
		reply_details.handle,
		already_liked
	ORDER BY
		tweets.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
