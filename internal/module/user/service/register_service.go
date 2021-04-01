package service

import (
	"time"

	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/HotPotatoC/twitter-clone/internal/token"
	"github.com/HotPotatoC/twitter-clone/pkg/bcrypt"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/HotPotatoC/twitter-clone/pkg/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type RegisterInput struct {
	Name     string `json:"name" validate:"required,alpha,excludesall= "`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (i RegisterInput) Validate() []*validator.ValidationError {
	return validator.ValidateStruct(i)
}

type RegisterService interface {
	Execute(input RegisterInput) (*token.AccessToken, *token.RefreshToken, error)
}

type registerService struct {
	db database.Database
}

func NewRegisterService(db database.Database) RegisterService {
	return registerService{db: db}
}

func (s registerService) Execute(input RegisterInput) (*token.AccessToken, *token.RefreshToken, error) {
	var alreadyExists bool

	err := s.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", input.Email).Scan(&alreadyExists)
	if err != nil {
		return nil, nil, errors.Wrap(err, "service.registerService.Execute")
	}

	if alreadyExists {
		return nil, nil, entity.ErrUserAlreadyExists
	}

	hash, err := bcrypt.Hash(input.Password)
	if err != nil {
		return nil, nil, errors.Wrap(err, "service.registerService.Execute")
	}

	_, err = s.db.Exec("INSERT INTO users(name, email, password, created_at) VALUES($1, $2, $3, $4)",
		input.Name, input.Email, hash, time.Now())
	if err != nil {
		return nil, nil, errors.Wrap(err, "service.registerService.Execute")
	}

	var id int64
	err = s.db.QueryRow("SELECT id FROM users WHERE email = $1", input.Email).Scan(&id)
	if err != nil {
		return nil, nil, errors.Wrap(err, "service.registerService.Execute")
	}

	at, err := token.NewAccessToken(jwt.MapClaims{
		"userID": id,
		"name":   input.Name,
		"email":  input.Email,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "service.registerService.Execute")
	}

	rt, err := token.NewRefreshToken(jwt.MapClaims{
		"userID": id,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "service.registerService.Execute")
	}


	return at, rt, nil
}
