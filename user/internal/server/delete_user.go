package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
	"github.com/jackc/pgx/v4"
	"github.com/twitchtv/twirp"
)

func (h *handler) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	if err := validateDeleteUserRequest(ctx, req); err != nil {
		return &user.DeleteUserResponse{Success: false}, err
	}

	err := h.service.DeleteUser(ctx, req.GetUserId())
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return &user.DeleteUserResponse{Success: false}, twirp.NotFoundError(fmt.Sprintf("User with id %s does not exists", req.GetUserId()))
		default:
			return &user.DeleteUserResponse{Success: false}, twirp.InternalErrorWith(err)
		}
	}

	return &user.DeleteUserResponse{Success: true}, nil
}

func validateDeleteUserRequest(ctx context.Context, req *user.DeleteUserRequest) error {
	if req.GetUserId() == "" {
		return twirp.RequiredArgumentError("user_id")
	}

	return nil
}
