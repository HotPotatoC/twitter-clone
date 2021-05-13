package entity

import "errors"

var (
	ErrTweetDoesNotExist     = errors.New("tweet does not exist")
	ErrTweetAlreadyFavorited = errors.New("tweet already favorited")
	ErrTweetAlreadyRetweeted = errors.New("tweet already retweeted")
)
