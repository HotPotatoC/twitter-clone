package server

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
)

func (h *handler) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	return nil, nil
}
