package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type GetUserOutput struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	FollowersCount  int       `json:"followers_count"`
	FollowingsCount int       `json:"followings_count"`
	JoinedAt        time.Time `json:"joined_at"`
}

type GetUserService interface {
	Execute(userID int64) (GetUserOutput, error)
}

type getUserService struct {
	db database.Database
}

func NewGetUserService(db database.Database) GetUserService {
	return getUserService{db: db}
}

func (s getUserService) Execute(userID int64) (GetUserOutput, error) {
	var userExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", userID).Scan(&userExists)
	if err != nil {
		return GetUserOutput{}, errors.Wrap(err, "service.getUserService.Execute")
	}

	if !userExists {
		return GetUserOutput{}, entity.ErrUserDoesNotExist
	}

	var id int64
	var name string
	var joinedAt time.Time
	var followingsCount, followersCount int
	err = s.db.QueryRow(`
	SELECT u.id,
		u.name,
		u.created_at,
		COUNT(f1.*) AS followings_count,
		COUNT(f2.*) AS followers_count
	FROM users AS u
		LEFT JOIN follows AS f1 ON f1.follower_id = u.id
		LEFT JOIN follows AS f2 ON f2.followed_id = u.id
	WHERE u.id = $1
	GROUP BY u.id
	`, userID).Scan(&id, &name, &joinedAt, &followingsCount, &followersCount)
	if err != nil {
		return GetUserOutput{}, errors.Wrap(err, "service.getUserService.Execute")
	}

	return GetUserOutput{
		ID:              id,
		Name:            name,
		FollowersCount:  followersCount,
		FollowingsCount: followingsCount,
		JoinedAt:        joinedAt,
	}, nil
}
