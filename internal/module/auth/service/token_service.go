package service

import (
	"errors"
	"fmt"

	"github.com/HotPotatoC/twitter-clone/internal/token"
	"github.com/HotPotatoC/twitter-clone/pkg/cache"
	"github.com/HotPotatoC/twitter-clone/pkg/config"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/dgrijalva/jwt-go"
	errorspkg "github.com/pkg/errors"
)

type TokenService interface {
	Execute(refreshToken string) (*token.AccessToken, error)
}

type tokenService struct {
	db    database.Database
	cache cache.Cache
}

func NewTokenService(db database.Database, cache cache.Cache) TokenService {
	return tokenService{
		db:    db,
		cache: cache,
	}
}

func (s tokenService) Execute(refreshToken string) (*token.AccessToken, error) {
	delimiter := config.GetString("REDIS_KEY_DELIMITER", "::")
	claims, err := token.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, errorspkg.Wrap(err, "service.tokenService.Execute")
	}

	key := fmt.Sprintf("ref_token%s%s", delimiter, claims["id"])

	// Check if refresh token is already blacklisted in the redis cache
	_, err = s.cache.Get(key)
	if err == nil {
		return nil, errors.New("refresh token is blacklisted")
	}

	var id int
	var name, email string
	err = s.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", claims["userID"]).Scan(&id, &name, &email)
	if err != nil {
		return nil, errorspkg.Wrap(err, "service.tokenService.Execute")
	}

	at, err := token.NewAccessToken(jwt.MapClaims{
		"userID": id,
		"name":   name,
		"email":  email,
	})
	if err != nil {
		return nil, errorspkg.Wrap(err, "service.tokenService.Execute")
	}

	return at, nil
}
