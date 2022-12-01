package server

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
)

func (h *handler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return nil, nil
}
