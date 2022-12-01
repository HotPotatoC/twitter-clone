package service

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/internal/models"
)

func (s *service) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	return models.User{}, nil
}
