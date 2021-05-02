package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/database"
	"github.com/HotPotatoC/twitter-clone/module/tweet/entity"
	"github.com/pkg/errors"
)

type ListTweetFeedOutput struct {
	entity.Tweet
	AuthorName     string        `json:"author_name"`
	AuthorHandle   string        `json:"author_handle"`
	AuthorPhotoURL string        `json:"author_photo_url"`
	Reply          *entity.Reply `json:"replied_to,omitempty"`
	IsReply        bool          `json:"is_reply"`
	AlreadyLiked   bool          `json:"already_liked"`
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
		var photoURLs []string
		var repliedToTweetAlreadyLiked sql.NullBool
		var repliedToTweetID, replyFavoriteCount, replyReplyCount sql.NullInt64
		var repliedToName, repliedToHandle, repliedToPhotoURL, replyContent sql.NullString
		var replyPhotoURLs []string
		var createdAt time.Time
		var favoritesCount, repliesCount int
		var alreadyLiked bool

		err = rows.Scan(&id, &content, &photoURLs, &createdAt, &name, &handle, &photoURL, &repliedToTweetID, &replyContent, &replyPhotoURLs, &repliedToName, &repliedToHandle, &repliedToPhotoURL, &repliedToTweetAlreadyLiked, &replyReplyCount, &replyFavoriteCount, &favoritesCount, &repliesCount, &alreadyLiked)
		if err != nil {
			return []ListTweetFeedOutput{}, errors.Wrap(err, "service.listTweetFeedService.Execute")
		}

		if repliedToTweetID.Valid {
			// The tweet is a reply
			tweets = append(tweets, ListTweetFeedOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					PhotoURLs:      photoURLs,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					CreatedAt:      createdAt,
				},
				AuthorName:     name,
				AuthorHandle:   handle,
				AuthorPhotoURL: photoURL,
				Reply: &entity.Reply{
					ID:             repliedToTweetID.Int64,
					Content:        replyContent.String,
					PhotoURLs:      replyPhotoURLs,
					AuthorName:     repliedToName.String,
					AuthorHandle:   repliedToHandle.String,
					AuthorPhotoURL: repliedToPhotoURL.String,
					FavoritesCount: int(replyFavoriteCount.Int64),
					RepliesCount:   int(replyReplyCount.Int64),
					AlreadyLiked:   repliedToTweetAlreadyLiked.Bool,
				},
				IsReply:      true,
				AlreadyLiked: alreadyLiked,
			})
		} else {
			tweets = append(tweets, ListTweetFeedOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					PhotoURLs:      photoURLs,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					CreatedAt:      createdAt,
				},
				AuthorName:     name,
				AuthorHandle:   handle,
				AuthorPhotoURL: photoURL,
				IsReply:        false,
				AlreadyLiked:   alreadyLiked,
			})
		}
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
		tweets.photo_urls,
		tweets.created_at,
		users.name,
		users.handle,
		users.photo_url,
		reply_details.id_tweet,
		reply_details.content,
		reply_details.photo_urls,
		reply_details.name,
		reply_details.handle,
		reply_details.photo_url,
		reply_details.already_liked,
		-- Reply's replies count
		(SELECT COUNT(replies.id_reply) FROM replies
			WHERE replies.id_tweet = reply_details.id_tweet),
		-- Reply's favorites count
		(SELECT COUNT(favorites.id) FROM favorites
			WHERE favorites.id_tweet = reply_details.id_tweet),
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
				tweets.content,
				tweets.photo_urls,
				users.name,
				users.handle,
				users.photo_url,
				CASE WHEN favorites.id_user = $1
					AND favorites.id_tweet = replies.id_tweet THEN
					TRUE
				ELSE
					FALSE
				END already_liked
			FROM replies
				INNER JOIN tweets ON tweets.id = replies.id_tweet
				INNER JOIN users ON users.id = tweets.id_user
				LEFT JOIN favorites ON favorites.id_tweet = tweets.id
			GROUP BY
				replies.id_reply,
				replies.id_tweet,
				tweets.content,
				tweets.photo_urls,
				users.name,
				users.handle,
				users.photo_url,
				favorites.id_user,
				favorites.id_tweet
			) AS reply_details ON reply_details.id_reply = tweets.id
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
		reply_details.content,
		reply_details.photo_urls,
		reply_details.name,
		reply_details.handle,
		reply_details.photo_url,
		reply_details.already_liked,
		favorites.id_user,
		favorites.id_tweet,
		already_liked
	ORDER BY
		tweets.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
