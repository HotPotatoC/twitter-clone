package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type ListTweetOutput struct {
	entity.Tweet
	Name string `json:"name"`
}

type ListTweetService interface {
	Execute() ([]ListTweetOutput, error)
}

type listTweetService struct {
	db database.Database
}

func NewListTweetService(db database.Database) ListTweetService {
	return listTweetService{db: db}
}

func (s listTweetService) Execute() ([]ListTweetOutput, error) {
	var tweets []ListTweetOutput
	rows, err := s.db.Query(`
	SELECT tweets.id,
		tweets.content,
		tweets.user_id,
		tweets.created_at,
		users.name
	FROM tweets
    INNER JOIN users ON tweets.user_id = users.id
	`)
	if err != nil {
		return []ListTweetOutput{}, errors.Wrap(err, "service.listTweetService.Execute")
	}
	defer rows.Close()

	for rows.Next() {
		var id, userID int64
		var content, name string
		var createdAt time.Time

		err = rows.Scan(&id, &content, &userID, &createdAt, &name)
		if err != nil {
			return []ListTweetOutput{}, errors.Wrap(err, "service.listTweetService.Execute")
		}

		tweets = append(tweets, ListTweetOutput{
			Tweet: entity.Tweet{
				ID:        id,
				Content:   content,
				CreatedAt: createdAt,
			},
			Name: name,
		})
	}

	if err := rows.Err(); err != nil {
		return []ListTweetOutput{}, errors.Wrap(err, "service.listTweetService.Execute")
	}

	return tweets, nil
}
