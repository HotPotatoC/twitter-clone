package service

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/clients"
	"github.com/HotPotatoC/twitter-clone/user/internal/models"
)

// Service is the interface that provides user related methods
type Service interface {
	// FindUserByID finds a user by id
	FindUserByID(ctx context.Context, id int64) (models.User, error)

	// FindUserByEmail finds a user by email
	FindUserByEmail(ctx context.Context, email string) (models.User, error)

	// CreateUser creates a new user
	CreateUser(ctx context.Context, params CreateUserParams) error

	// DeleteUser deletes an existing user
	DeleteUser(ctx context.Context, id int64) error
}

type service struct {
	clients clients.Clients
}

// NewService creates a new user business-layer service
func NewService(clients clients.Clients) Service {
	return &service{clients: clients}
}
