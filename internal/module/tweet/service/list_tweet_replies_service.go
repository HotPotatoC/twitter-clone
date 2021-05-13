package service

import (
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/pkg/errors"
)

type ListTweetRepliesOutput struct {
	entity.Tweet
	AuthorName     string `json:"author_name"`
	AuthorHandle   string `json:"author_handle"`
	AuthorPhotoURL string `json:"author_photo_url"`
	AlreadyLiked   bool   `json:"already_liked"`
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
			return []ListTweetRepliesOutput{}, module.ErrInvalidCursor
		}

		rows, err = s.db.Query(query, tweetID, cursor)
		if err != nil {
			return []ListTweetRepliesOutput{}, module.ErrInvalidCursor
		}
	} else {
		rows, err = s.db.Query(query, tweetID)
		if err != nil {
			return []ListTweetRepliesOutput{}, errors.Wrap(err, "service.listFollowersService.Execute")
		}
	}
	defer rows.Close()

	for rows.Next() {
		var id, authorID int64
		var favoritesCount, repliesCount int
		var content, authorName, authorHandle, authorPhotoURL string
		var photoURLs []string
		var createdAt time.Time
		var alreadyLiked bool

		err = rows.Scan(
			&id,
			&content,
			&photoURLs,
			&authorID,
			&createdAt,
			&authorName,
			&authorHandle,
			&authorPhotoURL,
			&favoritesCount,
			&repliesCount,
			&alreadyLiked)
		if err != nil {
			return []ListTweetRepliesOutput{}, errors.Wrap(err, "service.listTweetRepliesService.Execute")
		}

		tweets = append(tweets, ListTweetRepliesOutput{
			Tweet: entity.Tweet{
				ID:             id,
				Content:        content,
				PhotoURLs:      photoURLs,
				CreatedAt:      createdAt,
				FavoritesCount: favoritesCount,
				RepliesCount:   repliesCount,
			},
			AuthorName:     authorName,
			AuthorHandle:   authorHandle,
			AuthorPhotoURL: authorPhotoURL,
			AlreadyLiked:   alreadyLiked,
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
		tweets.photo_urls,
		tweets.id_user,
		tweets.created_at,
		users.name,
		users.handle,
		users.photo_url,
		COUNT(favorites.id),
		COUNT(r.id_reply),
        EXISTS (
            SELECT 1 FROM favorites
            WHERE favorites.id_tweet = tweets.id AND favorites.id_user = $1
		) AS already_liked
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
