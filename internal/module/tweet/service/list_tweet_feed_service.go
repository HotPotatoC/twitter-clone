package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type ListTweetFeedOutput struct {
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
	AlreadyLiked      bool   `json:"already_liked"`
}

type ListTweetFeedService interface {
	Execute(userID int64, createdAtCursor string) ([]ListTweetFeedOutput, error)
}

type listTweetFeedService struct {
	db database.Database
}

func NewListTweetFeedService(db database.Database) ListTweetFeedService {
	return listTweetFeedService{db: db}
}

func (s listTweetFeedService) Execute(userID int64, createdAtCursor string) ([]ListTweetFeedOutput, error) {
	var tweets []ListTweetFeedOutput

	var rows database.Rows
	var err error

	withCursor := createdAtCursor != ""
	query := s.buildSQLQuery(withCursor)

	if withCursor {
		cursor, err := time.Parse(time.RFC3339, createdAtCursor)
		if err != nil {
			return []ListTweetFeedOutput{}, ErrInvalidCursor
		}

		rows, err = s.db.Query(query, userID, cursor)
		if err != nil {
			return []ListTweetFeedOutput{}, errors.Wrap(err, "service.listTweetFeedService.Execute")
		}
	} else {
		rows, err = s.db.Query(query, userID)
		if err != nil {
			return []ListTweetFeedOutput{}, errors.Wrap(err, "service.listTweetFeedService.Execute")
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
		var alreadyLiked bool

		err = rows.Scan(&id, &content, &createdAt, &name, &handle, &photoURL, &repliedToTweetID, &repliedToName, &repliedToHandle, &repliedToPhotoURL, &favoritesCount, &repliesCount, &alreadyLiked)
		if err != nil {
			return []ListTweetFeedOutput{}, errors.Wrap(err, "service.listTweetFeedService.Execute")
		}

		tweets = append(tweets, ListTweetFeedOutput{
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
			AlreadyLiked:      alreadyLiked,
		})
	}

	if err := rows.Err(); err != nil {
		return []ListTweetFeedOutput{}, errors.Wrap(err, "service.listTweetFeedService.Execute")
	}

	return tweets, nil
}

func (s listTweetFeedService) buildSQLQuery(withCursor bool) string {
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
		COUNT(replies.id_reply),
		CASE WHEN favorites.id_user = $1
			AND favorites.id_tweet = tweets.id THEN
			TRUE
		ELSE
			FALSE
		END already_liked
	FROM tweets
		INNER JOIN users ON users.id = tweets.id_user
		INNER JOIN follows on follows.followed_id = users.id
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
	WHERE follows.follower_id = $1
	`)

	if withCursor {
		queryBuilder.WriteString("AND tweets.created_at < $2")
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
		reply_details.photo_url,
		already_liked
	ORDER BY
		tweets.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
