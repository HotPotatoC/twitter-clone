package service

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/clients"
	"github.com/HotPotatoC/twitter-clone/user/internal/models"
	"github.com/HotPotatoC/twitter-clone/user/internal/service/repository"
)

// Service is the interface that provides user related methods
type Service interface {
	// FindUserByID finds a user by id
	FindUserByID(ctx context.Context, id string) (models.User, error)

	// FindUserByEmail finds a user by email
	FindUserByEmail(ctx context.Context, email string) (models.User, error)

	// CreateUser creates a new user
	CreateUser(ctx context.Context, params CreateUserParams) (models.User, error)

	// DeleteUser deletes an existing user
	DeleteUser(ctx context.Context, id string) error
}

type service struct {
	clients    clients.Clients
	repository repository.Repository
}

// NewService creates a new user business-layer service
func NewService(clients clients.Clients) Service {
	return &service{
		clients:    clients,
		repository: repository.NewRepository(clients.WriterDB, clients.ReaderDB),
	}
}
