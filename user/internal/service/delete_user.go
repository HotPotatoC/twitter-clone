package service

import (
	"context"
)

func (s *service) DeleteUser(ctx context.Context, id string) error {
	return s.repository.DeleteUser(ctx, id)
}
