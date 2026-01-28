package main

import (
	"Hog-auth/internal/app/configs"
	"Hog-auth/internal/bootstrap/di"
	"Hog-auth/internal/pkg/fiber"
	"Hog-auth/internal/pkg/postgres"
	"Hog-auth/internal/pkg/redis"
	"Hog-auth/internal/pkg/validator"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		configs.Module,
		validator.Module,
		redis.Module,
		postgres.Module,
		fiber.Module,

		di.ProvidersModule,
		di.RepositoriesModule,
		di.ServicesModule,
		di.HandlersModule,
		di.RoutesModule,
	).Run()
}
