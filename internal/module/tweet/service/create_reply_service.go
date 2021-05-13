package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/common/validator"
	"github.com/HotPotatoC/twitter-clone/internal/module/tweet/entity"
	"github.com/pkg/errors"
)

type CreateReplyInput struct {
	Content string `json:"content" validate:"required,max=300"`
}

func (i CreateReplyInput) Validate() []*validator.ValidationError {
	return validator.ValidateStruct(i)
}

type CreateReplyService interface {
	Execute(input CreateReplyInput, userID int64, tweetID int64) error
}

type createReplyService struct {
	db database.Database
}

func NewCreateReplyService(db database.Database) CreateReplyService {
	return createReplyService{db: db}
}

func (s createReplyService) Execute(input CreateReplyInput, userID int64, tweetID int64) error {
	tx, err := s.db.BeginTx()
	if err != nil {
		return errors.Wrap(err, "service.createReplyService.Execute")
	}

	var tweetExists bool
	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM tweets WHERE id = $1)", tweetID).Scan(&tweetExists)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "service.createReplyService.Execute")
		}
		return errors.Wrap(err, "service.createReplyService.Execute")
	}

	if !tweetExists {
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "service.createReplyService.Execute")
		}
		return entity.ErrTweetDoesNotExist
	}

	var insertedReplyID int64
	err = tx.QueryRow("INSERT INTO tweets(content, id_user, created_at) VALUES($1, $2, $3) RETURNING id",
		input.Content, userID, time.Now()).Scan(&insertedReplyID)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "service.createReplyService.Execute")
		}
		return errors.Wrap(err, "service.createReplyService.Execute")
	}

	_, err = tx.Exec("INSERT INTO replies(id_reply, id_tweet) VALUES($1, $2)",
		insertedReplyID, tweetID)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return errors.Wrap(err, "service.createReplyService.Execute")
		}
		return errors.Wrap(err, "service.createReplyService.Execute")
	}

	return tx.Commit()
}
