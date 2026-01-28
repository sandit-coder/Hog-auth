package configs

import (
	"go.uber.org/fx"
)

var Module = fx.Module("configs",
	fx.Provide(
		NewFiber,
		NewPostgres,
		NewRedis,
	),
	fx.Invoke(
		LoadConfig,
	))
