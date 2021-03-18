package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/modules/user/entity"
	"github.com/HotPotatoC/twitter-clone/pkg/bcrypt"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/HotPotatoC/twitter-clone/pkg/validator"
	"github.com/pkg/errors"
)

type RegisterInput struct {
	Name     string `json:"name" validate:"required,alpha,excludesall= "`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" valudate:"required"`
}

func (i RegisterInput) Validate() []*validator.ValidationError {
	return validator.ValidateStruct(i)
}

type RegisterService interface {
	Execute(input RegisterInput) error
}

type registerService struct {
	db database.Database
}

func NewRegisterService(db database.Database) RegisterService {
	return registerService{db: db}
}

func (s registerService) Execute(input RegisterInput) error {
	tx, err := s.db.BeginTx()
	if err != nil {
		return err
	}

	var alreadyExists bool

	err = tx.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", input.Email).Scan(&alreadyExists)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "service.registerService.Execute")
	}

	if alreadyExists {
		tx.Rollback()
		return entity.ErrUserAlreadyExists
	}

	hash, err := bcrypt.Hash(input.Password)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "service.registerService.Execute")
	}

	_, err = tx.Exec("INSERT INTO users(name, email, password, created_at) VALUES($1, $2, $3, $4)",
		input.Name, input.Email, hash, time.Now())
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "service.registerService.Execute")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "service.registerService.Execute")
	}

	return nil
}
