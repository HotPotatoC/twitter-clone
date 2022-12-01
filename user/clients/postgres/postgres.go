package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgreSQLClient(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	connection, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := connection.Ping(ctx); err != nil {
		return nil, err
	}

	return connection, nil
}
