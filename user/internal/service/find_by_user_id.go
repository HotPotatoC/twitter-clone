package service

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/internal/models"
)

func (s *service) FindUserByID(ctx context.Context, id int64) (models.User, error) {
	sql := `SELECT
		id,
		name,
		screen_name,
		email,
		bio,
		location,
		website,
		birth_date,
		profile_image_url,
		profile_banner_url,
		followers_count,
		followings_count,
		created_at,
		updated_at
	FROM users WHERE id = $1`

	var user models.User

	err := s.clients.ReaderDB.QueryRow(ctx, sql, id).Scan(
		&user.ID,
		&user.Name,
		&user.ScreenName,
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
