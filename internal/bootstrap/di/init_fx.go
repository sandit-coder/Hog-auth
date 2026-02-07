package di

import (
	AuthHandlers "Hog-auth/internal/app/adapters/primary/http-adapter/handlers/auth"
	"Hog-auth/internal/app/adapters/primary/http-adapter/routes"
	"Hog-auth/internal/app/adapters/secondary/jwt"
	"Hog-auth/internal/app/adapters/secondary/kafka-adapter/kafka_client"
	"Hog-auth/internal/app/adapters/secondary/kafka-adapter/verification_code_publisher"
	"Hog-auth/internal/app/adapters/secondary/repositories/refresh_session"
	"Hog-auth/internal/app/adapters/secondary/repositories/user"
	"Hog-auth/internal/app/application/services/auth"

	"go.uber.org/fx"
)

var KafkaModule = fx.Module("kafka", fx.Provide(
	kafka_client.New,
	verification_code_publisher.New))

var ServicesModule = fx.Module("services", fx.Provide(
	auth.New,
	jwt.New))

var RepositoriesModule = fx.Module("repositories", fx.Provide(
	user.NewRepository,
	refresh_session.NewRepository,
))

var HandlersModule = fx.Module("handlers", fx.Provide(
	AuthHandlers.NewHandler))

var RoutesModule = fx.Module("routes", fx.Invoke(
	routes.AppendAuth,
))
