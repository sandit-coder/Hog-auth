package di

import (
	AuthHandlers "Hog-auth/internal/app/adapters/primary/http-adapter/handlers/auth"
	"Hog-auth/internal/app/adapters/primary/http-adapter/routes"
	"Hog-auth/internal/app/adapters/secondary/providers/cookie"
	"Hog-auth/internal/app/adapters/secondary/providers/jwt"
	"Hog-auth/internal/app/adapters/secondary/repositories/user"
	"Hog-auth/internal/app/application/services/auth"

	"go.uber.org/fx"
)

var ServicesModule = fx.Module("services", fx.Provide(
	auth.New))

var RepositoriesModule = fx.Module("repositories", fx.Provide(
	user.NewRepository,
))

var HandlersModule = fx.Module("handlers", fx.Provide(
	AuthHandlers.NewHandler))

var ProvidersModule = fx.Module("providers", fx.Provide(
	cookie.New,
	jwt.New))

var RoutesModule = fx.Module("routes", fx.Invoke(

	routes.AppendAuth,
))
