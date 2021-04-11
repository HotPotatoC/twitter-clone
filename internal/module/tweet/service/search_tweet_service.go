package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type SearchTweetOutput struct {
	entity.Tweet
	Name           string  `json:"name"`
	RepliedToTweet int64   `json:"replied_to_tweet_id,omitempty"`
	RepliedToName  string  `json:"replied_to_name,omitempty"`
	FavoritesCount int     `json:"favorites_count"`
	RepliesCount   int     `json:"replies_count"`
	Rank           float64 `json:"rank"`
}

type SearchTweetService interface {
	Execute(searchQuery string, cursor string) ([]SearchTweetOutput, error)
}

type searchTweetService struct {
	db database.Database
}

func NewSearchTweetService(db database.Database) SearchTweetService {
	return searchTweetService{db: db}
}

func (s searchTweetService) Execute(searchQuery string, cursor string) ([]SearchTweetOutput, error) {
	var tweets []SearchTweetOutput

	var rows database.Rows
	var err error

	withCursor := cursor != ""
	query := s.buildSQLQuery(withCursor)

	if withCursor {
		rows, err = s.db.Query(query, searchQuery, cursor)
		if err != nil {
			return []SearchTweetOutput{}, errors.Wrap(err, "service.searchTweetService.Execute")
		}
	} else {
		rows, err = s.db.Query(query, searchQuery)
		if err != nil {
			return []SearchTweetOutput{}, errors.Wrap(err, "service.searchTweetService.Execute")
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
		var rank float64

		err = rows.Scan(&id, &content, &createdAt, &name, &repliedToTweetID, &repliedToName, &favoritesCount, &repliesCount, &rank)
		if err != nil {
			return []SearchTweetOutput{}, errors.Wrap(err, "service.searchTweetService.Execute")
		}

		tweets = append(tweets, SearchTweetOutput{
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
			Rank:           rank,
		})
	}

	if err := rows.Err(); err != nil {
		return []SearchTweetOutput{}, errors.Wrap(err, "service.searchTweetService.Execute")
	}

	return tweets, nil
}

func (s searchTweetService) buildSQLQuery(withCursor bool) string {
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
		COUNT(r1.*),
		ts_rank(tweets.content_tsv, plainto_tsquery($1))
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
	WHERE tweets.content_tsv @@ plainto_tsquery($1)
	`)

	if withCursor {
		queryBuilder.WriteString("AND ts_rank(tweets.content_tsv, plainto_tsquery($1)) < $2")
	}

	queryBuilder.WriteString(`
	GROUP BY
		tweets.id
	ORDER BY
		ts_rank(tweets.content_tsv, plainto_tsquery($1)) DESC
	LIMIT 10`)

	return queryBuilder.String()
}
