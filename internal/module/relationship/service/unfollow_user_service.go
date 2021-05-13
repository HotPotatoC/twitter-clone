package service

import (
	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module/relationship/entity"
	userEntity "github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/pkg/errors"
)

type UnfollowUserService interface {
	Execute(followerID int64, followedID int64) (username string, err error)
}

type unfollowUserService struct {
	db database.Database
}

func NewUnfollowUserService(db database.Database) UnfollowUserService {
	return unfollowUserService{db: db}
}

func (s unfollowUserService) Execute(followerID int64, followedID int64) (string, error) {
	var userExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", followedID).Scan(&userExists)
	if err != nil {
		return "", errors.Wrap(err, "service.unfollowUserService.Execute")
	}

	if !userExists {
		return "", userEntity.ErrUserDoesNotExist
	}

	var userIsNotFollowing bool
	err = s.db.QueryRow(`
	SELECT NOT EXISTS (
        SELECT 1
        FROM follows
        WHERE follower_id = $1 AND followed_id = $2
    )`, followerID, followedID).Scan(&userIsNotFollowing)
	if err != nil {
		return "", errors.Wrap(err, "service.unfollowUserService.Execute")
	}

	if userIsNotFollowing {
		return "", entity.ErrUserIsNotFollowing
	}

	_, err = s.db.Exec("DELETE FROM follows WHERE follower_id = $1 AND followed_id = $2",
		followerID, followedID)
	if err != nil {
		return "", errors.Wrap(err, "service.unfollowUserService.Execute")
	}

	var name string
	err = s.db.QueryRow("SELECT name FROM users WHERE id = $1", followedID).Scan(&name)
	if err != nil {
		return "", errors.Wrap(err, "service.unfollowUserService.Execute")
	}

	return name, nil
}
