package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/HotPotatoC/twitter-clone/internal/token"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/pkg/errors"
)

type MeService interface {
	Execute(accessToken string) (*entity.User, error)
}

type meService struct {
	db database.Database
}

func NewMeService(db database.Database) MeService {
	return meService{db: db}
}

func (s meService) Execute(accessToken string) (*entity.User, error) {
	claims, err := token.VerifyAccessToken(accessToken)
	if err != nil {
		return nil, errors.Wrap(err, "service.meService.Execute")
	}

	var id int64
	var name, email string

	err = s.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", claims["userID"]).Scan(&id, &name, &email)

	if err != nil {
		return nil, errors.Wrap(err, "service.meService.Execute")
	}

	return entity.NewUser(id, name, email, "", time.Now()), nil
}
