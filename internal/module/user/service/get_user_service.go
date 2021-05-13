package service

import (
	"database/sql"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/pkg/errors"
)

type GetUserOutput struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Handle          string    `json:"handle"`
	Bio             string    `json:"bio"`
	Location        string    `json:"location"`
	Website         string    `json:"website"`
	BirthDate       time.Time `json:"birth_date"`
	FollowersCount  int       `json:"followers_count"`
	FollowingsCount int       `json:"followings_count"`
	IsFollowing     bool      `json:"is_following"`
	PhotoURL        string    `json:"photo_url"`
	JoinedAt        time.Time `json:"joined_at"`
}

type GetUserService interface {
	Execute(userID int64, username string) (GetUserOutput, error)
}

type getUserService struct {
	db database.Database
}

func NewGetUserService(db database.Database) GetUserService {
	return getUserService{db: db}
}

func (s getUserService) Execute(userID int64, username string) (GetUserOutput, error) {
	var userExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE handle = $1)", username).Scan(&userExists)
	if err != nil {
		return GetUserOutput{}, errors.Wrap(err, "service.getUserService.Execute")
	}

	if !userExists {
		return GetUserOutput{}, entity.ErrUserDoesNotExist
	}

	var id int64
	var name, handle, photoURL string
	var bio, location, website sql.NullString
	var birthDate sql.NullTime
	var joinedAt time.Time
	var followingsCount, followersCount int
	var isFollowing bool

	err = s.db.QueryRow(`
	SELECT users.id,
		users.name,
		users.handle,
		users.bio,
		users.location,
		users.website,
		users.birth_date,
		users.photo_url,
		users.created_at,
		COUNT(f1.*) AS followings_count,
		COUNT(f2.*) AS followers_count,
		CASE WHEN f2.follower_id = $1
			AND f2.followed_id = users.id THEN
			TRUE
		ELSE
			FALSE
		END is_following
	FROM users
		LEFT JOIN follows AS f1 ON f1.follower_id = users.id
		LEFT JOIN follows AS f2 ON f2.followed_id = users.id
	WHERE users.handle = $2
	GROUP BY
		users.id,
		is_following
	`, userID, username).Scan(&id, &name, &handle, &bio, &location, &website, &birthDate, &photoURL, &joinedAt, &followingsCount, &followersCount, &isFollowing)
	if err != nil {
		return GetUserOutput{}, errors.Wrap(err, "service.getUserService.Execute")
	}

	return GetUserOutput{
		ID:              id,
		Name:            name,
		Handle:          handle,
		PhotoURL:        photoURL,
		Bio:             bio.String,
		Location:        location.String,
		Website:         website.String,
		BirthDate:       birthDate.Time,
		FollowersCount:  followersCount,
		FollowingsCount: followingsCount,
		IsFollowing:     isFollowing,
		JoinedAt:        joinedAt,
	}, nil
}
