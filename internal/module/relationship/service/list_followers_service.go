package service

import (
	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/pkg/errors"
)

type ListFollowersOutput struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `json:"string"`
}

type ListFollowersService interface {
	Execute(userID int64) ([]ListFollowersOutput, error)
}

type listFollowersService struct {
	db database.Database
}

func NewListFollowersService(db database.Database) ListFollowersService {
	return listFollowersService{db: db}
}

func (s listFollowersService) Execute(userID int64) ([]ListFollowersOutput, error) {
	var userExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", userID).Scan(&userExists)
	if err != nil {
		return []ListFollowersOutput{}, errors.Wrap(err, "service.listFollowersService.Execute")
	}

	if !userExists {
		return []ListFollowersOutput{}, entity.ErrUserDoesNotExist
	}

	var followers []ListFollowersOutput
	rows, err := s.db.Query(`
	SELECT follows.id,
		follows.follower_id,
		users.id,
		users.name
	FROM follows
    INNER JOIN users ON follows.follower_id = users.id
	WHERE follows.followed_id = $1`, userID)
	if err != nil {
		return []ListFollowersOutput{}, errors.Wrap(err, "service.listFollowersService.Execute")
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var userID int64
		var name string

		err = rows.Scan(&id, nil, &userID, &name)
		if err != nil {
			return []ListFollowersOutput{}, errors.Wrap(err, "service.listFollowersService.Execute")
		}

		followers = append(followers, ListFollowersOutput{
			ID:     id,
			UserID: userID,
			Name:   name,
		})
	}

	if err := rows.Err(); err != nil {
		return []ListFollowersOutput{}, errors.Wrap(err, "service.listFollowersService.Execute")
	}

	return followers, nil
}
