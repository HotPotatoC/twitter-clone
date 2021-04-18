package service

import (
	"database/sql"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type GetUserOutput struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Bio             string    `json:"bio"`
	Location        string    `json:"location"`
	Website         string    `json:"website"`
	BirthDate       time.Time `json:"birth_date"`
	FollowersCount  int       `json:"followers_count"`
	FollowingsCount int       `json:"followings_count"`
	JoinedAt        time.Time `json:"joined_at"`
}

type GetUserService interface {
	Execute(username string) (GetUserOutput, error)
}

type getUserService struct {
	db database.Database
}

func NewGetUserService(db database.Database) GetUserService {
	return getUserService{db: db}
}

func (s getUserService) Execute(username string) (GetUserOutput, error) {
	var userExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE name = $1)", username).Scan(&userExists)
	if err != nil {
		return GetUserOutput{}, errors.Wrap(err, "service.getUserService.Execute")
	}

	if !userExists {
		return GetUserOutput{}, entity.ErrUserDoesNotExist
	}

	var id int64
	var name string
	var bio, location, website sql.NullString
	var birthDate sql.NullTime
	var joinedAt time.Time
	var followingsCount, followersCount int
	err = s.db.QueryRow(`
	SELECT u.id,
		u.name,
		u.bio,
		u.location,
		u.website,
		u.birth_date,
		u.created_at,
		COUNT(f1.*) AS followings_count,
		COUNT(f2.*) AS followers_count
	FROM users AS u
		LEFT JOIN follows AS f1 ON f1.follower_id = u.id
		LEFT JOIN follows AS f2 ON f2.followed_id = u.id
	WHERE u.name = $1
	GROUP BY u.id
	`, username).Scan(&id, &name, &bio, &location, &website, &birthDate, &joinedAt, &followingsCount, &followersCount)
	if err != nil {
		return GetUserOutput{}, errors.Wrap(err, "service.getUserService.Execute")
	}

	return GetUserOutput{
		ID:              id,
		Name:            name,
		Bio:             bio.String,
		Location:        location.String,
		Website:         website.String,
		BirthDate:       birthDate.Time,
		FollowersCount:  followersCount,
		FollowingsCount: followingsCount,
		JoinedAt:        joinedAt,
	}, nil
}
