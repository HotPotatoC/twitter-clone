package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/database"
	"github.com/HotPotatoC/twitter-clone/module/tweet/entity"
	"github.com/pkg/errors"
)

type GetUserTweetsOutput struct {
	entity.Tweet
	Name         string        `json:"author_name"`
	Handle       string        `json:"author_handle"`
	PhotoURL     string        `json:"author_photo_url"`
	Reply        *entity.Reply `json:"replied_to,omitempty"`
	IsReply      bool          `json:"is_reply"`
	AlreadyLiked bool          `json:"already_liked"`
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
		var content, name, handle, photoURL string
		var repliedToTweetAlreadyLiked sql.NullBool
		var repliedToTweetID, replyFavoriteCount, replyReplyCount sql.NullInt64
		var repliedToName, repliedToHandle, repliedToPhotoURL, replyContent sql.NullString
		var createdAt time.Time
		var favoritesCount, repliesCount int
		var alreadyLiked bool

		err = rows.Scan(&id, &content, &createdAt, &name, &handle, &photoURL, &repliedToTweetID, &replyContent, &repliedToName, &repliedToHandle, &repliedToPhotoURL, &repliedToTweetAlreadyLiked, &replyReplyCount, &replyFavoriteCount, &favoritesCount, &repliesCount, &alreadyLiked)
		if err != nil {
			return []GetUserTweetsOutput{}, errors.Wrap(err, "service.getUserTweetsService.Execute")
		}

		if repliedToTweetID.Valid {
			// The tweet is a reply
			tweets = append(tweets, GetUserTweetsOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					CreatedAt:      createdAt,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
				},
				Name:     name,
				Handle:   handle,
				PhotoURL: photoURL,
				Reply: &entity.Reply{
					ID:             repliedToTweetID.Int64,
					Content:        replyContent.String,
					AuthorName:     repliedToName.String,
					AuthorHandle:   repliedToHandle.String,
					AuthorPhotoURL: repliedToPhotoURL.String,
					RepliesCount:   int(replyReplyCount.Int64),
					FavoritesCount: int(replyFavoriteCount.Int64),
					AlreadyLiked:   repliedToTweetAlreadyLiked.Bool,
				},
				IsReply:      true,
				AlreadyLiked: alreadyLiked,
			})
		} else {
			tweets = append(tweets, GetUserTweetsOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					CreatedAt:      createdAt,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
				},
				Name:         name,
				Handle:       handle,
				PhotoURL:     photoURL,
				IsReply:      false,
				AlreadyLiked: alreadyLiked,
			})
		}
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
		users.photo_url,
		reply_details.id_tweet,
		reply_details.content,
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
	FROM
		tweets
		LEFT JOIN users ON users.id = tweets.id_user
		LEFT JOIN (
			SELECT
				replies.id_reply,
				replies.id_tweet,
				tweets.content,
				users.name,
				users.handle,
				users.photo_url,
				CASE WHEN favorites.id_user = $1
					AND favorites.id_tweet = replies.id_tweet THEN
					TRUE
				ELSE
					FALSE
				END already_liked
			FROM
				replies
				INNER JOIN tweets ON tweets.id = replies.id_tweet
				INNER JOIN users ON users.id = tweets.id_user
				LEFT JOIN favorites ON favorites.id_tweet = tweets.id
			GROUP BY
				replies.id_reply,
				replies.id_tweet,
				tweets.content,
				users.name,
				users.handle,
				users.photo_url,
				favorites.id_user,
				favorites.id_tweet
			) AS reply_details ON reply_details.id_reply = tweets.id
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
		users.photo_url,
		reply_details.id_tweet,
		reply_details.content,
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
