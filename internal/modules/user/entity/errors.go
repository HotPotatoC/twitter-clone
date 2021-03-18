package entity

import "errors"

var (
	ErrUserDoesNotExist  = errors.New("user does not exist")
	ErrUserAlreadyExists = errors.New("user already exists")
)
