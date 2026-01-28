package fiber

import "go.uber.org/fx"

var Module = fx.Module("fiber",
	fx.Provide(
		NewServer,
	), fx.Invoke(
		StartFiber,
	))
