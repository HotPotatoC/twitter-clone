package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/pkg/errors"
)

type FavoriteTweetService interface {
	Execute(tweetID int64, userID int64) error
}

type favoriteTweetService struct {
	db database.Database
}

func NewFavoriteTweetService(db database.Database) FavoriteTweetService {
	return favoriteTweetService{db: db}
}

func (s favoriteTweetService) Execute(tweetID int64, userID int64) error {
	var tweetExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM tweets WHERE id = $1)", tweetID).Scan(&tweetExists)
	if err != nil {
		return errors.Wrap(err, "service.favoriteTweetService.Execute")
	}

	if !tweetExists {
		return entity.ErrTweetDoesNotExist
	}

	var alreadyFavorited bool
	err = s.db.QueryRow(`
	SELECT EXISTS (
        SELECT 1
        FROM favorites
        WHERE id_tweet = $1 AND id_user = $2
    )`, tweetID, userID).Scan(&alreadyFavorited)
	if err != nil {
		return errors.Wrap(err, "service.favoriteTweetService.Execute")
	}

	if alreadyFavorited {
		_, err = s.db.Exec("DELETE FROM favorites WHERE id_tweet = $1 AND id_user = $2", tweetID, userID)
		if err != nil {
			return errors.Wrap(err, "service.favoriteTweetService.Execute")
		}
		return entity.ErrTweetAlreadyFavorited
	}

	_, err = s.db.Exec("INSERT INTO favorites(id_tweet, id_user, created_at) VALUES($1, $2, $3)",
		tweetID, userID, time.Now())
	if err != nil {
		return errors.Wrap(err, "service.favoriteTweetService.Execute")
	}

	return nil
}
