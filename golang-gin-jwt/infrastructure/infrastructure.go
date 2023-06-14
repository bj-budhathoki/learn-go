package infrastructure

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewRouter),
	// fx.Provide(NewDatabase),
	// fx.Provide(NewMigrations),
)
