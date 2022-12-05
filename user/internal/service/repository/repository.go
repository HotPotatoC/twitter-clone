package repository

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository is the interface that provides user related methods
type Repository interface {
	// FindUserByID finds a user by id
	FindUserByID(ctx context.Context, id string) (models.User, error)

	// FindUserByEmail finds a user by email
	FindUserByEmail(ctx context.Context, email string) (models.User, error)

	// CreateUser creates a new user
	CreateUser(ctx context.Context, params models.User) (models.User, error)

	// DeleteUser deletes an existing user
	DeleteUser(ctx context.Context, id string) error
}

type repository struct {
	writerDB *pgxpool.Pool
	readerDB *pgxpool.Pool

	queryBuilder squirrel.StatementBuilderType
}

func NewRepository(writerDB *pgxpool.Pool, readerDB *pgxpool.Pool) Repository {
	return &repository{
		writerDB:     writerDB,
		readerDB:     readerDB,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
