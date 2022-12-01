package server

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
)

func (h *handler) FindUserByEmail(ctx context.Context, req *user.FindUserByEmailRequest) (*user.FindUserByEmailResponse, error) {
	return nil, nil
}
