package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module/relationship/entity"
	userEntity "github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/pkg/errors"
)

type FollowUserService interface {
	Execute(followerID int64, followedID int64) (username string, err error)
}

type followUserService struct {
	db database.Database
}

func NewFollowUserService(db database.Database) FollowUserService {
	return followUserService{db: db}
}

func (s followUserService) Execute(followerID int64, followedID int64) (string, error) {
	var userExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", followedID).Scan(&userExists)
	if err != nil {
		return "", errors.Wrap(err, "service.followUserService.Execute")
	}

	if !userExists {
		return "", userEntity.ErrUserDoesNotExist
	}

	var userAlreadyFollowed bool
	err = s.db.QueryRow(`
	SELECT EXISTS (
        SELECT 1
        FROM follows
        WHERE follower_id = $1 AND followed_id = $2
    )`, followerID, followedID).Scan(&userAlreadyFollowed)
	if err != nil {
		return "", errors.Wrap(err, "service.followUserService.Execute")
	}

	if userAlreadyFollowed {
		return "", entity.ErrUserAlreadyFollowed
	}

	_, err = s.db.Exec("INSERT INTO follows(follower_id, followed_id, created_at) VALUES($1, $2, $3)",
		followerID, followedID, time.Now())
	if err != nil {
		return "", errors.Wrap(err, "service.followUserService.Execute")
	}

	var name string
	err = s.db.QueryRow("SELECT name FROM users WHERE id = $1", followedID).Scan(&name)
	if err != nil {
		return "", errors.Wrap(err, "service.followUserService.Execute")
	}

	return name, nil
}
