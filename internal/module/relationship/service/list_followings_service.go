package service

import (
	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/pkg/errors"
)

type ListFollowingsOutput struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `json:"string"`
}

type ListFollowingsService interface {
	Execute(userID int64) ([]ListFollowingsOutput, error)
}

type listFollowingsService struct {
	db database.Database
}

func NewListFollowingsService(db database.Database) ListFollowingsService {
	return listFollowingsService{db: db}
}

func (s listFollowingsService) Execute(userID int64) ([]ListFollowingsOutput, error) {
	var userExists bool
	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)", userID).Scan(&userExists)
	if err != nil {
		return []ListFollowingsOutput{}, errors.Wrap(err, "service.ListFollowingsService.Execute")
	}

	if !userExists {
		return []ListFollowingsOutput{}, entity.ErrUserDoesNotExist
	}

	var followings []ListFollowingsOutput
	rows, err := s.db.Query(`
	SELECT follows.id,
		follows.followed_id,
		users.id,
		users.name
	FROM follows
    INNER JOIN users ON follows.followed_id = users.id
	WHERE follows.follower_id = $1`, userID)
	if err != nil {
		return []ListFollowingsOutput{}, errors.Wrap(err, "service.ListFollowingsService.Execute")
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var userID int64
		var name string

		err = rows.Scan(&id, nil, &userID, &name)
		if err != nil {
			return []ListFollowingsOutput{}, errors.Wrap(err, "service.ListFollowingsService.Execute")
		}

		followings = append(followings, ListFollowingsOutput{
			ID:     id,
			UserID: userID,
			Name:   name,
		})
	}

	if err := rows.Err(); err != nil {
		return []ListFollowingsOutput{}, errors.Wrap(err, "service.ListFollowingsService.Execute")
	}

	return followings, nil
}
