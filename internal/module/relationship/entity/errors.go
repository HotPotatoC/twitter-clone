package entity

import "errors"

var (
	ErrUserAlreadyFollowed = errors.New("user already followed")
	ErrUserIsNotFollowing  = errors.New("user is not following")
)
