package service

import (
	"github.com/HotPotatoC/twitter-clone/internal/common/bcrypt"
	"github.com/HotPotatoC/twitter-clone/internal/common/database"
	"github.com/HotPotatoC/twitter-clone/internal/common/token"
	"github.com/HotPotatoC/twitter-clone/internal/common/validator"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var (
	ErrInvalidPassword = errors.New("Invalid password provided")
)

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (i LoginInput) Validate() []*validator.ValidationError {
	return validator.ValidateStruct(i)
}

type LoginService interface {
	Execute(input LoginInput) (*token.AccessToken, *token.RefreshToken, error)
}

type loginService struct {
	db database.Database
}

func NewLoginService(db database.Database) LoginService {
	return loginService{db: db}
}

func (s loginService) Execute(input LoginInput) (*token.AccessToken, *token.RefreshToken, error) {
	var id int
	var handle, email, password string
	err := s.db.QueryRow("SELECT id, handle, email, password FROM users WHERE email = $1", input.Email).Scan(&id, &handle, &email, &password)
	if err != nil {
		return nil, nil, entity.ErrUserDoesNotExist
	}

	if !bcrypt.Compare(password, input.Password) {
		return nil, nil, ErrInvalidPassword
	}

	at, err := token.NewAccessToken(jwt.MapClaims{
		"userID": id,
		"handle": handle,
		"email":  email,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "service.loginService.Execute")
	}

	rt, err := token.NewRefreshToken(jwt.MapClaims{
		"userID": id,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "service.loginService.Execute")
	}

	return at, rt, nil
}
