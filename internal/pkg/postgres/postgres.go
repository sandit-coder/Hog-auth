package postgres

import (
	"context"

	trmpg "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func New(cfg *Config) *pgxpool.Pool {
	connectionString := GetPostgresConnectionString(cfg.Dsn)

	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil
	}

	return db
}

func NewTransactionManager(pool *pgxpool.Pool) trm.Manager {
	return manager.Must(trmpg.NewDefaultFactory(pool))
}

func StartPostgres(lc fx.Lifecycle, pool *pgxpool.Pool) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := pool.Ping(ctx); err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			pool.Close()
			return nil
		},
	})
}
