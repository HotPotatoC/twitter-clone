package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/pkg/errors"
)

type ListUserTweetsOutput struct {
	entity.Tweet
	AuthorName          string        `json:"author_name"`
	AuthorHandle        string        `json:"author_handle"`
	AuthorPhotoURL      string        `json:"author_photo_url"`
	Reply               *entity.Reply `json:"replied_to,omitempty"`
	IsReply             bool          `json:"is_reply"`
	AlreadyLiked        bool          `json:"already_liked"`
	IsRetweet           bool          `json:"is_retweet"`
	RetweetAuthorHandle string        `json:"retweet_author_handle"`
	AlreadyRetweeted    bool          `json:"already_retweeted"`
}

type ListUserTweetsService interface {
	Execute(userID int64, username string, createdAtCursor string) ([]ListUserTweetsOutput, error)
}

type listUserTweetsService struct {
	db database.Database
}

func NewListUserTweetsService(db database.Database) ListUserTweetsService {
	return listUserTweetsService{db: db}
}

func (s listUserTweetsService) Execute(userID int64, username string, createdAtCursor string) ([]ListUserTweetsOutput, error) {
	var tweets []ListUserTweetsOutput

	var rows database.Rows
	var err error

	withCursor := createdAtCursor != ""
	query := s.buildSQLQuery(withCursor)

	if withCursor {
		cursor, err := time.Parse(time.RFC3339, createdAtCursor)
		if err != nil {
			return []ListUserTweetsOutput{}, module.ErrInvalidCursor
		}

		rows, err = s.db.Query(query, userID, username, cursor)
		if err != nil {
			return []ListUserTweetsOutput{}, errors.Wrap(err, "service.listUserTweetsService.Execute")
		}
	} else {
		rows, err = s.db.Query(query, userID, username)
		if err != nil {
			return []ListUserTweetsOutput{}, errors.Wrap(err, "service.listUserTweetsService.Execute")
		}
	}
	defer rows.Close()

	for rows.Next() {
		var id, authorID int64
		var content, authorName, authorHandle, authorPhotoURL string
		var photoURLs []string
		var repliedTweetAlreadyLiked sql.NullBool
		var repliedTweetID, repliedTweetFavoriteCount, repliedTweetReplyCount sql.NullInt64
		var repliedTweetAuthorName, repliedTweetAuthorHandle, repliedTweetAuthorPhotoURL, replyContent, retweetAuthorHandle sql.NullString
		var replyPhotoURLs []string
		var createdAt time.Time
		var favoritesCount, repliesCount, retweetsCount int
		var alreadyLiked, alreadyRetweeted, isRetweet bool

		err = rows.Scan(
			&id,
			&content,
			&photoURLs,
			&createdAt,
			&authorID,
			&authorName,
			&authorHandle,
			&authorPhotoURL,
			&alreadyLiked,
			&favoritesCount,
			&repliesCount,
			&repliedTweetID,
			&replyContent,
			&replyPhotoURLs,
			&repliedTweetAuthorName,
			&repliedTweetAuthorHandle,
			&repliedTweetAuthorPhotoURL,
			&repliedTweetAlreadyLiked,
			&repliedTweetReplyCount,
			&repliedTweetFavoriteCount,
			&isRetweet,
			&retweetAuthorHandle,
			&alreadyRetweeted,
			&retweetsCount)
		if err != nil {
			return []ListUserTweetsOutput{}, errors.Wrap(err, "service.listUserTweetsService.Execute")
		}

		switch {
		case isRetweet:
			tweets = append(tweets, ListUserTweetsOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					PhotoURLs:      photoURLs,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					RetweetsCount:  retweetsCount,
					CreatedAt:      createdAt,
				},
				AuthorName:          authorName,
				AuthorHandle:        authorHandle,
				AuthorPhotoURL:      authorPhotoURL,
				IsRetweet:           true,
				RetweetAuthorHandle: retweetAuthorHandle.String,
				AlreadyLiked:        alreadyLiked,
				AlreadyRetweeted:    alreadyRetweeted,
			})
		case repliedTweetID.Valid:
			// The tweet is a reply
			tweets = append(tweets, ListUserTweetsOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					PhotoURLs:      photoURLs,
					CreatedAt:      createdAt,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					RetweetsCount:  retweetsCount,
				},
				AuthorName:     authorName,
				AuthorHandle:   authorHandle,
				AuthorPhotoURL: authorPhotoURL,
				Reply: &entity.Reply{
					ID:             repliedTweetID.Int64,
					Content:        replyContent.String,
					PhotoURLs:      replyPhotoURLs,
					AuthorName:     repliedTweetAuthorName.String,
					AuthorHandle:   repliedTweetAuthorHandle.String,
					AuthorPhotoURL: repliedTweetAuthorPhotoURL.String,
					FavoritesCount: int(repliedTweetFavoriteCount.Int64),
					RepliesCount:   int(repliedTweetReplyCount.Int64),
					RetweetsCount:  retweetsCount,
					AlreadyLiked:   repliedTweetAlreadyLiked.Bool,
				},
				IsReply:          true,
				AlreadyLiked:     alreadyLiked,
				AlreadyRetweeted: alreadyRetweeted,
			})
		default:
			tweets = append(tweets, ListUserTweetsOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					PhotoURLs:      photoURLs,
					CreatedAt:      createdAt,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					RetweetsCount:  retweetsCount,
				},
				AuthorName:       authorName,
				AuthorHandle:     authorHandle,
				AuthorPhotoURL:   authorPhotoURL,
				IsReply:          false,
				AlreadyLiked:     alreadyLiked,
				AlreadyRetweeted: alreadyRetweeted,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return []ListUserTweetsOutput{}, errors.Wrap(err, "service.listUserTweetsService.Execute")
	}

	return tweets, nil
}

func (s listUserTweetsService) buildSQLQuery(withCursor bool) string {
	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
	WITH __tweets AS (
		SELECT
			tweets.id,
			tweets.content,
			tweets.photo_urls,
			tweets.created_at,
			users.id AS author_id,
			users.name AS author_name,
			users.handle AS author_handle,
			users.photo_url AS author_photo_url,
			COALESCE(COUNT(DISTINCT replies.id_reply), 0) AS reply_count,
			COALESCE(COUNT(DISTINCT favorites.id), 0) AS favorite_count,
			EXISTS (
				SELECT 1 FROM favorites
				WHERE favorites.id_tweet = tweets.id AND favorites.id_user = $1
			) AS already_liked
			FROM tweets
				INNER JOIN users ON tweets.id_user = users.id
				LEFT JOIN favorites ON tweets.id = favorites.id_tweet
				LEFT JOIN replies ON tweets.id = replies.id_tweet
			GROUP BY
				tweets.id,
				tweets.content,
				tweets.photo_urls,
				tweets.created_at,
				author_id,
				author_name,
				author_handle,
				author_photo_url
	)
	SELECT * FROM (
		-- User's tweets and replies
		SELECT
			__tweets.id,
			__tweets.content,
			__tweets.photo_urls,
			__tweets.created_at,
			__tweets.author_id,
			__tweets.author_name,
			__tweets.author_handle,
			__tweets.author_photo_url,
			__tweets.already_liked,
			__tweets.favorite_count,
			__tweets.reply_count,
			__replied_tweet.id_tweet AS replied_tweet_id_tweet,
			__replied_tweet.content AS replied_tweet_content,
			__replied_tweet.photo_urls AS replied_tweet_photo_urls,
			__replied_tweet.author_name AS replied_tweet_author_name,
			__replied_tweet.author_handle AS replied_tweet_author_handle,
			__replied_tweet.author_photo_url AS replied_tweet_author_photo_url,
			__replied_tweet.already_liked AS replied_tweet_already_liked,
			__replied_tweet.reply_count AS replied_tweet_reply_count,
			__replied_tweet.favorite_count AS replied_tweet_favorite_count,
			FALSE AS is_retweet,
			'' AS retweet_author_handle,
			FALSE AS already_retweeted,
			0 AS retweets_count
		FROM __tweets
			LEFT JOIN (
				SELECT
					replies.id_reply,
					replies.id_tweet,
					__tweets.content,
					__tweets.photo_urls,
					__tweets.author_name,
					__tweets.author_handle,
					__tweets.author_photo_url,
					__tweets.already_liked,
					__tweets.reply_count,
					__tweets.favorite_count
				FROM replies
					INNER JOIN __tweets ON replies.id_tweet = __tweets.id
			) AS __replied_tweet ON __tweets.id = __replied_tweet.id_reply
		WHERE __tweets.author_handle = $2
		UNION
		-- User retweets
		SELECT
			__tweets.id,
			__tweets.content,
			__tweets.photo_urls,
			retweets.created_at,
			__tweets.author_id,
			__tweets.author_name,
			__tweets.author_handle,
			__tweets.author_photo_url,
			__tweets.already_liked,
			__tweets.favorite_count,
			__tweets.reply_count,
			CAST(NULL AS int) AS replied_tweet_id_tweet,
			CAST(NULL AS varchar) AS replied_tweet_content,
			CAST(NULL AS text[]) AS replied_tweet_photo_urls,
			CAST(NULL AS varchar) AS replied_tweet_author_name,
			CAST(NULL AS varchar) AS replied_tweet_author_handle,
			CAST(NULL AS text) AS replied_tweet_author_photo_url,
			CAST(NULL AS boolean) AS replied_tweet_already_liked,
			CAST(NULL AS int) AS replied_tweet_reply_count,
			CAST(NULL AS int) AS replied_tweet_favorite_count,
			TRUE AS is_retweet,
			__retweet_author.handle AS retweet_author_handle,
            EXISTS (
                SELECT 1 FROM retweets
                WHERE retweets.id_user = $1 AND retweets.id_tweet = __tweets.id
			) AS already_retweeted,
			COUNT(retweets.id_tweet) AS retweets_count
		FROM  __tweets
			INNER JOIN retweets ON __tweets.id = retweets.id_tweet
			INNER JOIN users AS __retweet_author ON retweets.id_user = __retweet_author.id
		WHERE retweets.id_user = $1
		GROUP BY
			__tweets.id,
			__tweets.content,
			__tweets.photo_urls,
			retweets.created_at,
			__tweets.author_id,
			__tweets.author_name,
			__tweets.author_handle,
			__tweets.author_photo_url,
			__tweets.already_liked,
			__tweets.favorite_count,
			__tweets.reply_count,
			is_retweet,
			__retweet_author.handle,
			already_retweeted
	) __result
	`)

	if withCursor {
		queryBuilder.WriteString("WHERE __result.created_at < $3")
	}

	queryBuilder.WriteString(`
	ORDER BY
		__result.created_at DESC
	LIMIT 10`)

	return queryBuilder.String()
}
