package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

type Config struct {
	Dsn string `koanf:"dsn"`
}

func New(cfg Config) *redis.Client {
	opt, err := redis.ParseURL(cfg.Dsn)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	return client
}

func StartRedis(lc fx.Lifecycle, client *redis.Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := client.Ping(ctx).Err(); err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})
}
