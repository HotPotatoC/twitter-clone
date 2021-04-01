package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type ListTweetRepliesOutput struct {
	entity.Tweet
	Name string `json:"name"`
}

type ListTweetRepliesService interface {
	Execute(tweetID int64) ([]ListTweetRepliesOutput, error)
}

type listTweetRepliesService struct {
	db database.Database
}

func NewListTweetRepliesService(db database.Database) ListTweetRepliesService {
	return listTweetRepliesService{db: db}
}

func (s listTweetRepliesService) Execute(tweetID int64) ([]ListTweetRepliesOutput, error) {
	var tweetExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM tweets WHERE id = $1)", tweetID).Scan(&tweetExists)
	if err != nil {
		return nil, errors.Wrap(err, "service.listFollowersService.Execute")
	}

	if !tweetExists {
		return nil, entity.ErrTweetDoesNotExist
	}

	var tweets []ListTweetRepliesOutput
	rows, err := s.db.Query(`
	SELECT t.id,
		t.content,
		t.id_user,
		t.created_at,
		u.name
	FROM replies as r
	INNER JOIN tweets as t ON r.id_reply = t.id
    INNER JOIN users as u ON t.id_user = u.id
	WHERE r.id_tweet = $1
	`, tweetID)
	if err != nil {
		return []ListTweetRepliesOutput{}, errors.Wrap(err, "service.listTweetRepliesService.Execute")
	}
	defer rows.Close()

	for rows.Next() {
		var id, userID int64
		var content, name string
		var createdAt time.Time

		err = rows.Scan(&id, &content, &userID, &createdAt, &name)
		if err != nil {
			return []ListTweetRepliesOutput{}, errors.Wrap(err, "service.listTweetRepliesService.Execute")
		}

		tweets = append(tweets, ListTweetRepliesOutput{
			Tweet: entity.Tweet{
				ID:        id,
				Content:   content,
				CreatedAt: createdAt,
			},
			Name: name,
		})
	}

	if err := rows.Err(); err != nil {
		return []ListTweetRepliesOutput{}, errors.Wrap(err, "service.listTweetRepliesService.Execute")
	}

	return tweets, nil
}
