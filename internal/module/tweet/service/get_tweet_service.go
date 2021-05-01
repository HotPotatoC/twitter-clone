package service

import (
	"database/sql"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type GetTweetOutput struct {
	entity.Tweet
	Name         string        `json:"author_name"`
	Handle       string        `json:"author_handle"`
	PhotoURL     string        `json:"author_photo_url"`
	Reply        *entity.Reply `json:"replied_to,omitempty"`
	IsReply      bool          `json:"is_reply"`
	AlreadyLiked bool          `json:"already_liked"`
}

type GetTweetService interface {
	Execute(userID, tweetID int64) (GetTweetOutput, error)
}

type getTweetService struct {
	db database.Database
}

func NewGetTweetService(db database.Database) GetTweetService {
	return getTweetService{db: db}
}

func (s getTweetService) Execute(userID, tweetID int64) (GetTweetOutput, error) {
	var tweetExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM tweets WHERE id = $1)", tweetID).Scan(&tweetExists)
	if err != nil {
		return GetTweetOutput{}, errors.Wrap(err, "service.favoriteTweetService.Execute")
	}

	if !tweetExists {
		return GetTweetOutput{}, entity.ErrTweetDoesNotExist
	}

	var id int64
	var content, name, handle, photoURL string
	var repliedToTweetAlreadyLiked sql.NullBool
	var repliedToTweetID, replyFavoriteCount, replyReplyCount sql.NullInt64
	var repliedToName, repliedToHandle, repliedToPhotoURL, replyContent sql.NullString
	var createdAt time.Time
	var favoritesCount, repliesCount int
	var alreadyLiked bool

	err = s.db.QueryRow(`
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
			WHERE replies.id_tweet = tweets.id),
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
			FROM replies
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
		) as reply_details ON reply_details.id_reply = tweets.id
		LEFT JOIN favorites ON favorites.id_tweet = tweets.id
		LEFT JOIN replies ON replies.id_tweet = tweets.id
	WHERE tweets.id = $2
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
	`, userID, tweetID).Scan(&id, &content, &createdAt, &name, &handle, &photoURL, &repliedToTweetID, &replyContent, &repliedToName, &repliedToHandle, &repliedToPhotoURL, &repliedToTweetAlreadyLiked, &replyFavoriteCount, &replyReplyCount, &favoritesCount, &repliesCount, &alreadyLiked)
	if err != nil {
		return GetTweetOutput{}, errors.Wrap(err, "service.getTweetService.Execute")
	}

	if repliedToTweetID.Valid {
		// The tweet is a reply
		return GetTweetOutput{
			Tweet: entity.Tweet{
				ID:             id,
				Content:        content,
				FavoritesCount: favoritesCount,
				RepliesCount:   favoritesCount,
				CreatedAt:      createdAt,
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
				FavoritesCount: int(replyFavoriteCount.Int64),
				RepliesCount:   int(replyReplyCount.Int64),
				AlreadyLiked:   repliedToTweetAlreadyLiked.Bool,
			},
			IsReply:      true,
			AlreadyLiked: alreadyLiked,
		}, nil
	}

	return GetTweetOutput{
		Tweet: entity.Tweet{
			ID:             id,
			Content:        content,
			FavoritesCount: favoritesCount,
			RepliesCount:   favoritesCount,
			CreatedAt:      createdAt,
		},
		Name:         name,
		Handle:       handle,
		PhotoURL:     photoURL,
		IsReply:      false,
		AlreadyLiked: alreadyLiked,
	}, nil
}
