package server

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/internal/service"
	"github.com/HotPotatoC/twitter-clone/user/rpc/user"
	"github.com/twitchtv/twirp"
)

func (h *handler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	if err := validateCreateUserRequest(ctx, req); err != nil {
		return nil, err
	}

	_, err := h.service.FindUserByEmail(ctx, req.GetEmail())
	if err == nil {
		return nil, twirp.NewError(twirp.AlreadyExists, "email already exists")
	}

	createdUser, err := h.service.CreateUser(ctx, service.CreateUserParams{
		Name:             req.GetName(),
		ScreenName:       req.GetScreenName(),
		RawPassword:      req.GetPassword(),
		Email:            req.GetEmail(),
		Bio:              req.GetBio(),
		Location:         req.GetLocation(),
		Website:          req.GetWebsite(),
		ProfileImageURL:  req.GetProfileImageUrl(),
		ProfileBannerURL: req.GetProfileBannerUrl(),
		BirthDate:        req.GetBirthDate().AsTime(),
	})
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &user.CreateUserResponse{
		User: createdUser.PB(),
	}, nil
}

func validateCreateUserRequest(ctx context.Context, req *user.CreateUserRequest) error {
	if req.GetName() == "" {
		return twirp.RequiredArgumentError("name")
	}

	if req.GetScreenName() == "" {
		return twirp.RequiredArgumentError("screen_name")
	}

	if req.GetPassword() == "" {
		return twirp.RequiredArgumentError("email")
	}

	if req.GetEmail() == "" {
		return twirp.RequiredArgumentError("email")
	}

	return nil
}
