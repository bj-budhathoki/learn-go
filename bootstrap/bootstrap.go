package bootstrap

import (
	"context"

	"github.com/bj-budhathoki/learn-go/infrastructure"
	"go.uber.org/fx"
)

var Module = fx.Options(
	infrastructure.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, handler infrastructure.Router, env infrastructure.Env, logger infrastructure.Logger, database infrastructure.Database) {
	appStop := func(context.Context) error {
		logger.Zap.Info("Stapping application")
		conn, _ := database.DB.DB()
		conn.Close()
		return nil
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("------------------------")
			logger.Zap.Info("------ Working 📺 ------")
			logger.Zap.Info("------------------------")

			// logger.Zap.Info("Migrating DB schema...")
			go func() {

				if env.SERVER_PORT == "" {
					handler.Gin.Run(":5000")
				} else {
					handler.Gin.Run(":" + env.SERVER_PORT)
				}
			}()
			return nil
		},
		OnStop: appStop,
	})

}
