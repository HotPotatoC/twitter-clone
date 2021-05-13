package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/common/validator"
	"github.com/pkg/errors"
)

type UpdateUserInput struct {
	DisplayName string `json:"display_name" form:"display_name" validate:"omitempty,max=255"`
	Bio         string `json:"bio" form:"bio" validate:"omitempty,max=255"`
	Location    string `json:"location" form:"location" validate:"omitempty,max=30"`
	Website     string `json:"website" form:"website" validate:"omitempty,url"`
	BirthDate   string `json:"birth_date" form:"birth_date" validate:"omitempty,datetime=2006-01-02"`
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

func (s updateUserService) updateWithoutBirthDate(input UpdateUserInput, userID int64) error {
	_, err := s.db.Exec(`
		UPDATE users
		SET name = CASE WHEN name <> $1 THEN $1 ELSE $1 END,
			bio = CASE WHEN bio <> $2 THEN $2 ELSE $2 END,
			location = CASE WHEN location <> $3 THEN $3 ELSE $3 END,
			website = CASE WHEN website <> $4 THEN $4 ELSE $4 END
		WHERE id = $5`, input.DisplayName, input.Bio, input.Location, input.Website, userID)
	if err != nil {
		return errors.Wrap(err, "service.updateUserService.Execute")
	}

	return nil
}

func (s updateUserService) updateWithBirthDate(input UpdateUserInput, userID int64) error {
	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		return errors.Wrap(err, "service.updateUserService.Execute")
	}

	_, err = s.db.Exec(`
		UPDATE users
		SET name = CASE WHEN name <> $1 THEN $1 ELSE $1 END,
			bio = CASE WHEN bio <> $2 THEN $2 ELSE $2 END,
			location = CASE WHEN location <> $3 THEN $3 ELSE $3 END,
			website = CASE WHEN website <> $4 THEN $4 ELSE $4 END,
			birth_date = CASE WHEN birth_date <> $5 THEN $5 ELSE $5 END
		WHERE id = $6`, input.DisplayName, input.Bio, input.Location, input.Website, birthDate, userID)
	if err != nil {
		return errors.Wrap(err, "service.updateUserService.Execute")
	}

	return nil
}
