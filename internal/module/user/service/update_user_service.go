package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/HotPotatoC/twitter-clone/pkg/validator"
	"github.com/pkg/errors"
)

type UpdateUserInput struct {
	Bio       string `json:"bio" validate:"omitempty,max=255"`
	Location  string `json:"location" validate:"omitempty,max=30"`
	Website   string `json:"website" validate:"omitempty,url"`
	BirthDate string `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
}

func (i UpdateUserInput) Validate() []*validator.ValidationError {
	return validator.ValidateStruct(i)
}

type UpdateUserService interface {
	Execute(input UpdateUserInput, userID int64) error
}

type updateUserService struct {
	db database.Database
}

func NewUpdateUserService(db database.Database) UpdateUserService {
	return updateUserService{db: db}
}

func (s updateUserService) Execute(input UpdateUserInput, userID int64) error {
	if input.BirthDate == "" {
		return s.updateWithoutBirthDate(input, userID)
	}
	return s.updateWithBirthDate(input, userID)
}

func (s updateUserService) updateWithBirthDate(input UpdateUserInput, userID int64) error {
	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return errors.Wrap(err, "service.updateUserService.Execute")
	}

	_, err = s.db.Exec(`
		UPDATE users
		SET bio = CASE WHEN bio <> $1 THEN $1 ELSE $1 END,
			location = CASE WHEN location <> $2 THEN $2 ELSE $2 END,
			website = CASE WHEN website <> $3 THEN $3 ELSE $3 END,
			birth_date = CASE WHEN birth_date <> $4 THEN $4 ELSE $4 END
		WHERE id = $5`, input.Bio, input.Location, input.Website, birthDate, userID)
	if err != nil {
		return errors.Wrap(err, "service.updateUserService.Execute")
	}

	return nil
}

func (s updateUserService) updateWithoutBirthDate(input UpdateUserInput, userID int64) error {
	_, err := s.db.Exec(`
		UPDATE users
		SET bio = CASE WHEN bio <> $1 THEN $1 ELSE $1 END,
			location = CASE WHEN location <> $2 THEN $2 ELSE $2 END,
			website = CASE WHEN website <> $3 THEN $3 ELSE $3 END
		WHERE id = $4`, input.Bio, input.Location, input.Website, userID)
	if err != nil {
		return errors.Wrap(err, "service.updateUserService.Execute")
	}

	return nil
}
