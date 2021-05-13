package service

import (
	"fmt"
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/cache"
	"github.com/HotPotatoC/twitter-clone/internal/common/config"
	"github.com/HotPotatoC/twitter-clone/internal/common/token"
	"github.com/pkg/errors"
)

type LogoutService interface {
	Execute(refreshToken string) error
}

type logoutService struct {
	cache cache.Cache
}

func NewLogoutService(cache cache.Cache) LogoutService {
	return logoutService{
		cache: cache,
	}
}

func (s logoutService) Execute(refreshToken string) error {
	delimiter := config.GetString("REDIS_KEY_DELIMITER", "::")
	claims, err := token.VerifyRefreshToken(refreshToken)
	if err != nil {
		return errors.Wrap(err, "service.logoutService.Execute")
	}

	tokenID := claims["id"].(string)
	exp := claims["exp"].(float64) * float64(time.Microsecond)
	key := fmt.Sprintf("ref_token%s%s", delimiter, tokenID)

	err = s.cache.Set(key, refreshToken, time.Duration(exp))
	if err != nil {
		return errors.Wrap(err, "service.logoutService.Execute")
	}
	return nil
}
