package service

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/internal/models"
)

func (s *service) FindUserByID(ctx context.Context, id string) (models.User, error) {
	return s.repository.FindUserByID(ctx, id)
}
