package configs

import (
	"Hog-auth/internal/app/adapters/secondary/jwt"
	"Hog-auth/internal/app/adapters/secondary/kafka-adapter/kafka_client"
	"Hog-auth/internal/pkg/fiber"
	"Hog-auth/internal/pkg/postgres"
	"Hog-auth/internal/pkg/redis"

	"go.uber.org/fx"
)

var Module = fx.Module("configs",
	fx.Provide(
		Load,
		Redis,
		Postgres,
		Fiber,
		Jwt,
		Kafka,
	),
)

func Redis(full Config) redis.Config {
	return full.Infrastructure.Databases.Redis
}

func Postgres(full Config) postgres.Config {
	return full.Infrastructure.Databases.Postgres
}

func Jwt(full Config) jwt.Config {
	return full.Infrastructure.Jwt
}

func Fiber(full Config) fiber.Config {
	return full.Adapter.Primary.Fiber
}

func Kafka(full Config) kafka_client.Config {
	return full.Adapter.Secondary.KafkaClient
}
