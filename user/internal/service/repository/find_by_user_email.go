package repository

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/internal/models"
	"github.com/Masterminds/squirrel"
)

func (r *repository) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	query, args, _ := r.queryBuilder.
		Select("*").
		From("users").
		Where(squirrel.Eq{"email": email}).
		ToSql()

	var user models.User

	err := r.readerDB.QueryRow(ctx, query, args...).Scan(
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
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
