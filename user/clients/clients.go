package clients

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/clients/postgres"
	"github.com/HotPotatoC/twitter-clone/user/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/sync/errgroup"
)

type Clients struct {
	WriterDB *pgxpool.Pool
	ReaderDB *pgxpool.Pool
}

func NewClients(ctx context.Context, cfg *config.Config) (Clients, error) {
	var group errgroup.Group

	c := Clients{}

	group.Go(func() error {
		var err error
		c.WriterDB, err = postgres.NewPostgreSQLClient(ctx, cfg.Clients.WriterDbURL)
		if err != nil {
			return err
		}

		return nil
	})

	group.Go(func() error {
		var err error
		c.ReaderDB, err = postgres.NewPostgreSQLClient(ctx, cfg.Clients.ReaderDbURL)
		if err != nil {
			return err
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		return Clients{}, err
	}

	return c, nil
}
