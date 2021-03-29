package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/HotPotatoC/twitter-clone/pkg/validator"
	"github.com/pkg/errors"
)

type CreateTweetInput struct {
	Content string `json:"content" validate:"required,max=300"`
}

func (i CreateTweetInput) Validate() []*validator.ValidationError {
	return validator.ValidateStruct(i)
}

type CreateTweetService interface {
	Execute(input CreateTweetInput, userID int64) error
}

type createTweetService struct {
	db database.Database
}

func NewCreateTweetService(db database.Database) CreateTweetService {
	return createTweetService{db: db}
}

func (s createTweetService) Execute(input CreateTweetInput, userID int64) error {
	_, err := s.db.Exec("INSERT INTO tweets(content, id_user, created_at) VALUES($1, $2, $3)",
		input.Content, userID, time.Now())
	if err != nil {
		return errors.Wrap(err, "service.createTweetService.Execute")
	}

	return nil
}
