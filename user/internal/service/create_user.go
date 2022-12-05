package service

import (
	"context"
	"time"

	"github.com/HotPotatoC/twitter-clone/user/internal/models"
	"github.com/google/uuid"
)

type CreateUserParams struct {
	Name             string    `json:"name"`
	ScreenName       string    `json:"screen_name"`
	RawPassword      string    `json:"password"`
	Email            string    `json:"email"`
	Bio              string    `json:"bio"`
	Location         string    `json:"location"`
	Website          string    `json:"website"`
	ProfileImageURL  string    `json:"profile_image_url"`
	ProfileBannerURL string    `json:"profile_banner_url"`
	BirthDate        time.Time `json:"birth_date"`
}

func (s *service) CreateUser(ctx context.Context, params CreateUserParams) (models.User, error) {
	password := models.Password(params.RawPassword)

	if err := password.Validate(); err != nil {
		return models.User{}, err
	}

	hash, err := password.GenerateHash()
	if err != nil {
		return models.User{}, err
	}

	user, err := s.repository.CreateUser(ctx, models.User{
		ID:               uuid.New().String(),
		Name:             params.Name,
		ScreenName:       params.ScreenName,
		PasswordHash:     hash,
		Email:            params.Email,
		Bio:              params.Bio,
		Location:         params.Location,
		Website:          params.Website,
		ProfileImageURL:  params.ProfileImageURL,
		ProfileBannerURL: params.ProfileBannerURL,
		BirthDate:        params.BirthDate,
		FollowersCount:   0,
		FollowingsCount:  0,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
