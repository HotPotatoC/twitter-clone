package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type ListTweetOutput struct {
	entity.Tweet
	Name           string `json:"name"`
	RepliedToTweet int64  `json:"replied_to_tweet_id,omitempty"`
	RepliedToName  string `json:"replied_to_name,omitempty"`
	FavoritesCount int    `json:"favorites_count"`
	RepliesCount   int    `json:"replies_count"`
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
			return []ListTweetOutput{}, ErrInvalidCursor
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
		var content, name string
		var repliedToTweetID sql.NullInt64
		var repliedToName sql.NullString
		var createdAt time.Time
		var favoritesCount, repliesCount int

		err = rows.Scan(&id, &content, &createdAt, &name, &repliedToTweetID, &repliedToName, &favoritesCount, &repliesCount)
		if err != nil {
			return []ListTweetOutput{}, errors.Wrap(err, "service.listTweetService.Execute")
		}

		tweets = append(tweets, ListTweetOutput{
			Tweet: entity.Tweet{
				ID:        id,
				Content:   content,
				CreatedAt: createdAt,
			},
			Name:           name,
			RepliedToTweet: repliedToTweetID.Int64,
			RepliedToName:  repliedToName.String,
			FavoritesCount: favoritesCount,
			RepliesCount:   repliesCount,
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
		(ARRAY_AGG(users.name))[1],
		(ARRAY_AGG(sq.id_tweet))[1],
		(ARRAY_AGG(sq.name))[1],
		COUNT(f.*),
		COUNT(r1.*)
	FROM
		tweets
		LEFT JOIN users ON users.id = tweets.id_user
		LEFT JOIN (
			SELECT
				replies.id_reply,
				replies.id_tweet,
				users.name
			FROM
				replies
				INNER JOIN tweets AS t ON t.id = replies.id_tweet
				INNER JOIN users ON users.id = t.id_user) AS sq ON sq.id_reply = tweets.id
		LEFT JOIN favorites AS f ON f.id_tweet = tweets.id
		LEFT JOIN replies AS r1 ON r1.id_tweet = tweets.id
	`)

	if withCursor {
		queryBuilder.WriteString("WHERE tweets.created_at < $1")
	}

	queryBuilder.WriteString(`
	GROUP BY
		tweets.id
	ORDER BY
		tweets.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
