package repository

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/internal/models"
)

type CreateUserParams struct{}

func (r *repository) CreateUser(ctx context.Context, params models.User) (models.User, error) {
	input := map[string]any{
		"id":                 params.ID,
		"name":               params.Name,
		"screen_name":        params.ScreenName,
		"password_hash":      params.PasswordHash,
		"email":              params.Email,
		"bio":                params.Bio,
		"location":           params.Location,
		"website":            params.Website,
		"profile_image_url":  params.ProfileImageURL,
		"profile_banner_url": params.ProfileBannerURL,
		"birth_date":         params.BirthDate,
		"followers_count":    params.FollowersCount,
		"followings_count":   params.FollowingsCount,
		"created_at":         params.CreatedAt,
		"updated_at":         params.UpdatedAt,
	}

	query, args, err := r.queryBuilder.
		Insert("users").
		SetMap(input).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	var user models.User

	err = r.writerDB.QueryRow(ctx, query, args...).Scan(
		&user.ID,
		&user.Name,
		&user.ScreenName,
		&user.PasswordHash,
		&user.Email,
		&user.Bio,
		&user.Location,
		&user.Website,
		&user.BirthDate,
		&user.ProfileImageURL,
		&user.ProfileBannerURL,
		&user.FollowersCount,
		&user.FollowingsCount,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
