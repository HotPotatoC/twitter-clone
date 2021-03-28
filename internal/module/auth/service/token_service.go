package service

import (
	"github.com/HotPotatoC/twitter-clone/internal/token"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type TokenService interface {
	Execute(refreshToken string) (*token.AccessToken, error)
}

type tokenService struct {
	db database.Database
}

func NewTokenService(db database.Database) TokenService {
	return tokenService{db: db}
}

func (s tokenService) Execute(refreshToken string) (*token.AccessToken, error) {
	claims, err := token.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.Wrap(err, "service.refreshTokenService.Execute")
	}

	var id int
	var name, email string
	err = s.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", claims["userID"]).Scan(&id, &name, &email)
	if err != nil {
		return nil, errors.Wrap(err, "service.refreshTokenService.Execute")
	}

	at, err := token.NewAccessToken(jwt.MapClaims{
		"userID": id,
		"name":   name,
		"email":  email,
	})
	if err != nil {
		return nil, errors.Wrap(err, "service.refreshTokenService.Execute")
	}

	return at, nil
}
