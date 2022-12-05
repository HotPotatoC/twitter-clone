package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
	"github.com/jackc/pgx/v4"
	"github.com/twitchtv/twirp"
)

func (h *handler) FindUserByID(ctx context.Context, req *user.FindUserByIDRequest) (*user.FindUserByIDResponse, error) {
	if err := validateFindUserByIDRequest(ctx, req); err != nil {
		return nil, err
	}

	userData, err := h.service.FindUserByID(ctx, req.GetUserId())
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, twirp.NotFoundError(fmt.Sprintf("User with id %s does not exists", req.GetUserId()))
		default:
			return nil, twirp.InternalErrorWith(err)
		}
	}

	return &user.FindUserByIDResponse{
		User: userData.PB(),
	}, nil
}

func validateFindUserByIDRequest(ctx context.Context, req *user.FindUserByIDRequest) error {
	if req.GetUserId() == "" {
		return twirp.RequiredArgumentError("user_id")
	}

	return nil
}
