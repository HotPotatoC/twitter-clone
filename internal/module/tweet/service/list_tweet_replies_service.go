package service

import (
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type ListTweetRepliesOutput struct {
	entity.Tweet
	Name           string `json:"name"`
	FavoritesCount int    `json:"favorites_count"`
	RepliesCount   int    `json:"replies_count"`
}

type ListTweetRepliesService interface {
	Execute(tweetID int64, createdAtCursor string) ([]ListTweetRepliesOutput, error)
}

type listTweetRepliesService struct {
	db database.Database
}

func NewListTweetRepliesService(db database.Database) ListTweetRepliesService {
	return listTweetRepliesService{db: db}
}

func (s listTweetRepliesService) Execute(tweetID int64, createdAtCursor string) ([]ListTweetRepliesOutput, error) {
	var tweetExists bool
	var rows database.Rows
	var err error

	err = s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM tweets WHERE id = $1)", tweetID).Scan(&tweetExists)
	if err != nil {
		return nil, errors.Wrap(err, "service.listFollowersService.Execute")
	}

	if !tweetExists {
		return nil, entity.ErrTweetDoesNotExist
	}

	var tweets []ListTweetRepliesOutput

	withCursor := createdAtCursor != ""
	query := s.buildSQLQuery(withCursor)

	if withCursor {
		cursor, err := time.Parse(time.RFC3339, createdAtCursor)
		if err != nil {
			return []ListTweetRepliesOutput{}, ErrInvalidCursor
		}

		rows, err = s.db.Query(query, tweetID, cursor)
		if err != nil {
			return []ListTweetRepliesOutput{}, ErrInvalidCursor
		}
	} else {
		rows, err = s.db.Query(query, tweetID)
		if err != nil {
			return []ListTweetRepliesOutput{}, ErrInvalidCursor
		}
	}
	defer rows.Close()

	for rows.Next() {
		var id, userID int64
		var favoritesCount, repliesCount int
		var content, name string
		var createdAt time.Time

		err = rows.Scan(&id, &content, &userID, &createdAt, &name, &favoritesCount, &repliesCount)
		if err != nil {
			return []ListTweetRepliesOutput{}, errors.Wrap(err, "service.listTweetRepliesService.Execute")
		}

		tweets = append(tweets, ListTweetRepliesOutput{
			Tweet: entity.Tweet{
				ID:        id,
				Content:   content,
				CreatedAt: createdAt,
			},
			Name:           name,
			FavoritesCount: favoritesCount,
			RepliesCount:   repliesCount,
		})
	}

	if err := rows.Err(); err != nil {
		return []ListTweetRepliesOutput{}, errors.Wrap(err, "service.listTweetRepliesService.Execute")
	}

	return tweets, nil
}

func (s listTweetRepliesService) buildSQLQuery(withCursor bool) string {
	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
	SELECT
		t.id,
		t.content,
		t.id_user,
		t.created_at,
		(array_agg(u.name)) [1],
		COUNT(f.*),
		COUNT(r.*)
	FROM
		replies
		INNER JOIN tweets AS t ON t.id = replies.id_reply
		INNER JOIN users AS u ON t.id_user = u.id
		LEFT JOIN favorites AS f ON f.id_tweet = replies.id_reply
		LEFT JOIN replies as r ON r.id_tweet = t.id
	WHERE
		replies.id_tweet = $1 `)

	if withCursor {
		queryBuilder.WriteString("AND t.created_at < $2")
	}

	queryBuilder.WriteString(`
	GROUP BY
		t.id
	ORDER BY
		t.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
