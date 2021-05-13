package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/pkg/errors"
)

type SearchTweetOutput struct {
	entity.Tweet
	AuthorName     string        `json:"author_name"`
	AuthorHandle   string        `json:"author_handle"`
	AuthorPhotoURL string        `json:"author_photo_url"`
	Reply          *entity.Reply `json:"replied_to,omitempty"`
	IsReply        bool          `json:"is_reply"`
	Rank           float64       `json:"rank"`
	AlreadyLiked   bool          `json:"already_liked"`
}

type SearchTweetService interface {
	Execute(searchQuery string, userID int64, cursor string) ([]SearchTweetOutput, error)
}

type searchTweetService struct {
	db database.Database
}

func NewSearchTweetService(db database.Database) SearchTweetService {
	return searchTweetService{db: db}
}

func (s searchTweetService) Execute(searchQuery string, userID int64, cursor string) ([]SearchTweetOutput, error) {
	var tweets []SearchTweetOutput

	var rows database.Rows
	var err error

	withCursor := cursor != ""
	query := s.buildSQLQuery(withCursor)

	if withCursor {
		rows, err = s.db.Query(query, searchQuery, userID, cursor)
		if err != nil {
			return []SearchTweetOutput{}, errors.Wrap(err, "service.searchTweetService.Execute")
		}
	} else {
		rows, err = s.db.Query(query, searchQuery, userID)
		if err != nil {
			return []SearchTweetOutput{}, errors.Wrap(err, "service.searchTweetService.Execute")
		}
	}

	defer rows.Close()

	for rows.Next() {
		var id, authorID int64
		var content, authorName, authorHandle, authorPhotoURL string
		var rank float64
		var photoURLs []string
		var repliedTweetAlreadyLiked sql.NullBool
		var repliedTweetID, repliedTweetFavoriteCount, repliedTweetReplyCount sql.NullInt64
		var repliedTweetAuthorName, repliedTweetAuthorHandle, repliedTweetAuthorPhotoURL, replyContent sql.NullString
		var replyPhotoURLs []string
		var createdAt time.Time
		var favoritesCount, repliesCount int
		var alreadyLiked bool

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
			&rank,
			&repliedTweetID,
			&replyContent,
			&replyPhotoURLs,
			&repliedTweetAuthorName,
			&repliedTweetAuthorHandle,
			&repliedTweetAuthorPhotoURL,
			&repliedTweetAlreadyLiked,
			&repliedTweetReplyCount,
			&repliedTweetFavoriteCount)
		if err != nil {
			return []SearchTweetOutput{}, errors.Wrap(err, "service.searchTweetService.Execute")
		}

		if repliedTweetID.Valid {
			// The tweet is a reply
			tweets = append(tweets, SearchTweetOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					PhotoURLs:      photoURLs,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					CreatedAt:      createdAt,
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
					AlreadyLiked:   repliedTweetAlreadyLiked.Bool,
				},
				IsReply:      true,
				AlreadyLiked: alreadyLiked,
				Rank:         rank,
			})
		} else {
			tweets = append(tweets, SearchTweetOutput{
				Tweet: entity.Tweet{
					ID:             id,
					Content:        content,
					PhotoURLs:      photoURLs,
					FavoritesCount: favoritesCount,
					RepliesCount:   repliesCount,
					CreatedAt:      createdAt,
				},
				AuthorName:     authorName,
				AuthorHandle:   authorHandle,
				AuthorPhotoURL: authorPhotoURL,
				IsReply:        false,
				AlreadyLiked:   alreadyLiked,
				Rank:           rank,
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

	queryBuilder.WriteString(`WITH __tweets AS (
		SELECT
			tweets.id,
			tweets.content,
			tweets.photo_urls,
			tweets.created_at,
			tweets.content_tsv,
			ts_rank(tweets.content_tsv, plainto_tsquery($1)) AS rank,
			users.id AS author_id,
			users.name AS author_name,
			users.handle AS author_handle,
			users.photo_url AS author_photo_url,
			COALESCE(COUNT(DISTINCT replies.id_reply), 0) AS reply_count,
			COALESCE(COUNT(DISTINCT favorites.id), 0) AS favorite_count,
			EXISTS (
				SELECT
					1
				FROM
					favorites
				WHERE
					favorites.id_tweet = tweets.id
					AND favorites.id_user = $2) AS already_liked
			FROM
				tweets
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
		__tweets.rank,
		__replied_tweet.id_tweet,
		__replied_tweet.content,
		__replied_tweet.photo_urls,
		__replied_tweet.author_name,
		__replied_tweet.author_handle,
		__replied_tweet.author_photo_url,
		__replied_tweet.already_liked,
		__replied_tweet.reply_count,
		__replied_tweet.favorite_count
	FROM
		__tweets
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
			FROM
				replies
				INNER JOIN __tweets ON replies.id_tweet = __tweets.id
		) AS __replied_tweet ON __tweets.id = __replied_tweet.id_reply
	WHERE
		__tweets.content_tsv @@ plainto_tsquery($1)
	`)

	if withCursor {
		queryBuilder.WriteString("AND __tweets.rank < $3")
	}

	queryBuilder.WriteString(`
	ORDER BY
		__tweets.rank DESC
	LIMIT 10`)

	return queryBuilder.String()
}
