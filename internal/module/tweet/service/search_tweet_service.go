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
	Name     string        `json:"author_name"`
	Handle   string        `json:"author_handle"`
	PhotoURL string        `json:"author_photo_url"`
	Reply    *entity.Reply `json:"replied_to,omitempty"`
	IsReply  bool          `json:"is_reply"`
	Rank     float64       `json:"rank"`
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
		var content, name, handle, photoURL string
		var repliedToTweetID, replyFavoriteCount, replyReplyCount sql.NullInt64
		var repliedToName, repliedToHandle, repliedToPhotoURL, replyContent sql.NullString
		var createdAt time.Time
		var favoritesCount, repliesCount int
		var rank float64

		err = rows.Scan(&id, &content, &createdAt, &name, &handle, &photoURL, &repliedToTweetID, &replyContent, &repliedToName, &repliedToHandle, &repliedToPhotoURL, &replyFavoriteCount, &replyReplyCount, &favoritesCount, &repliesCount, &rank)
		if err != nil {
			return []SearchTweetOutput{}, errors.Wrap(err, "service.searchTweetService.Execute")
		}

		if repliedToTweetID.Valid {
			// The tweet is a reply
			tweets = append(tweets, SearchTweetOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					CreatedAt:      createdAt,
				},
				Name:     name,
				Handle:   handle,
				PhotoURL: photoURL,
				Reply: &entity.Reply{
					ID:             repliedToTweetID.Int64,
					AuthorName:     repliedToName.String,
					AuthorHandle:   repliedToHandle.String,
					AuthorPhotoURL: repliedToPhotoURL.String,
				},
				IsReply: true,
				Rank:    rank,
			})
		} else {
			tweets = append(tweets, SearchTweetOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					CreatedAt:      createdAt,
				},
				Name:     name,
				Handle:   handle,
				PhotoURL: photoURL,
				Rank:     rank,
			})
		}
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
		users.name,
		users.handle,
		users.photo_url,
		reply_details.id_tweet,
		reply_details.content,
		reply_details.name,
		reply_details.handle,
		reply_details.photo_url,
		-- Reply's replies count
		(SELECT COUNT(replies.id_reply) FROM replies
			WHERE replies.id_tweet = reply_details.id_tweet),
		-- Reply's favorites count
		(SELECT COUNT(favorites.id) FROM favorites
			WHERE favorites.id_tweet = reply_details.id_tweet),
		COUNT(favorites.id),
		COUNT(replies.id_reply),
		ts_rank(tweets.content_tsv, plainto_tsquery($1))
	FROM tweets
		LEFT JOIN users ON users.id = tweets.id_user
		LEFT JOIN (
			SELECT
				replies.id_reply,
				replies.id_tweet,
				t.content,
				users.name,
				users.handle,
				users.photo_url
			FROM replies
				INNER JOIN tweets AS t ON t.id = replies.id_tweet
				INNER JOIN users ON users.id = t.id_user) AS reply_details ON reply_details.id_reply = tweets.id
		LEFT JOIN favorites ON favorites.id_tweet = tweets.id
		LEFT JOIN replies ON replies.id_tweet = tweets.id
	WHERE tweets.content_tsv @@ plainto_tsquery($1)
	`)

	if withCursor {
		queryBuilder.WriteString("AND ts_rank(tweets.content_tsv, plainto_tsquery($1)) < $2")
	}

	queryBuilder.WriteString(`
	GROUP BY
		tweets.id,
		users.name,
		users.handle,
		users.photo_url,
		reply_details.id_tweet,
		reply_details.content,
		reply_details.name,
		reply_details.handle,
		reply_details.photo_url
	ORDER BY
		ts_rank(tweets.content_tsv, plainto_tsquery($1)) DESC
	LIMIT 10`)

	return queryBuilder.String()
}
