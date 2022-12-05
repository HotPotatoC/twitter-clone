package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
	"github.com/jackc/pgx/v4"
	"github.com/twitchtv/twirp"
)

func (h *handler) FindUserByEmail(ctx context.Context, req *user.FindUserByEmailRequest) (*user.FindUserByEmailResponse, error) {
	if err := validateFindUserByEmailRequest(ctx, req); err != nil {
		return nil, err
	}

	userData, err := h.service.FindUserByEmail(ctx, req.GetEmail())
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, twirp.NotFoundError(fmt.Sprintf("User with email %s does not exists", req.GetEmail()))
		default:
			return nil, twirp.InternalErrorWith(err)
		}
	}

	return &user.FindUserByEmailResponse{
		User: userData.PB(),
	}, nil
}

func validateFindUserByEmailRequest(ctx context.Context, req *user.FindUserByEmailRequest) error {
	if req.GetEmail() == "" {
		return twirp.RequiredArgumentError("email")
	}

	return nil
}
