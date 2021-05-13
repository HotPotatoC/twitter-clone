package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/pkg/errors"
)

type RetweetService interface {
	Execute(tweetID int64, userID int64) error
}

type retweetService struct {
	db database.Database
}

func NewRetweetService(db database.Database) RetweetService {
	return retweetService{db: db}
}

func (s retweetService) Execute(tweetID int64, userID int64) error {
	var tweetExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM tweets WHERE id = $1)", tweetID).Scan(&tweetExists)
	if err != nil {
		return errors.Wrap(err, "service.retweetService.Execute")
	}

	if !tweetExists {
		return entity.ErrTweetDoesNotExist
	}

	var alreadyRetweeted bool
	err = s.db.QueryRow(`
	SELECT EXISTS (
        SELECT 1
        FROM retweets
        WHERE id_tweet = $1 AND id_user = $2
    )`, tweetID, userID).Scan(&alreadyRetweeted)
	if err != nil {
		return errors.Wrap(err, "service.retweetService.Execute")
	}

	if alreadyRetweeted {
		_, err = s.db.Exec("DELETE FROM retweets WHERE id_tweet = $1 AND id_user = $2", tweetID, userID)
		if err != nil {
			return errors.Wrap(err, "service.retweetService.Execute")
		}
		return entity.ErrTweetAlreadyRetweeted
	}

	_, err = s.db.Exec("INSERT INTO retweets(id_tweet, id_user, created_at) VALUES($1, $2, $3)",
		tweetID, userID, time.Now())
	if err != nil {
		return errors.Wrap(err, "service.retweetService.Execute")
	}

	return nil
}
