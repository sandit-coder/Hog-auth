package configs

import (
	"Hog-auth/internal/app/adapters/secondary/jwt"
	"Hog-auth/internal/pkg/postgres"
	"Hog-auth/internal/pkg/redis"
)

type Config struct {
	Adapter        Adapters       `koanf:"adapter"`
	Infrastructure Infrastructure `koanf:"infrastructure"`
}

type Infrastructure struct {
	Jwt       jwt.Config `koanf:"jwt"`
	Databases Databases  `koanf:"databases"`
}

type Databases struct {
	Postgres postgres.Config `koanf:"postgres"`
	Redis    redis.Config    `koanf:"redis"`
}
