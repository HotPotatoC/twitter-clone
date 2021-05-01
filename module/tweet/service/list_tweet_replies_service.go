package service

import (
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/database"
	"github.com/HotPotatoC/twitter-clone/module/tweet/entity"
	"github.com/pkg/errors"
)

type ListTweetRepliesOutput struct {
	entity.Tweet
	Name           string `json:"author_name"`
	Handle         string `json:"author_handle"`
	PhotoURL       string `json:"author_photo_url"`
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
			return []ListTweetRepliesOutput{}, errors.Wrap(err, "service.listFollowersService.Execute")
		}
	}
	defer rows.Close()

	for rows.Next() {
		var id, userID int64
		var favoritesCount, repliesCount int
		var content, name, handle, photoURL string
		var createdAt time.Time

		err = rows.Scan(&id, &content, &userID, &createdAt, &name, &handle, &photoURL, &favoritesCount, &repliesCount)
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
			Handle:         handle,
			PhotoURL:       photoURL,
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
		tweets.id,
		tweets.content,
		tweets.id_user,
		tweets.created_at,
		users.name,
		users.handle,
		users.photo_url,
		COUNT(favorites.id),
		COUNT(r.id_reply)
	FROM replies
		INNER JOIN tweets ON tweets.id = replies.id_reply
		INNER JOIN users ON users.id = tweets.id_user
		LEFT JOIN favorites ON favorites.id_tweet = replies.id_reply
		LEFT JOIN replies as r ON r.id_tweet = tweets.id
	WHERE replies.id_tweet = $1
	`)

	if withCursor {
		queryBuilder.WriteString("AND tweets.created_at < $2")
	}

	queryBuilder.WriteString(`
	GROUP BY
		tweets.id,
		users.name,
		users.handle,
		users.photo_url
	ORDER BY
		tweets.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
