package bootstrap

import (
	"context"
	"fmt"

	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/api/controllers"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/api/routes"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/infrastructure"
	"go.uber.org/fx"
)

var Module = fx.Options(
	infrastructure.Module,
	routes.Module,
	controllers.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, handler infrastructure.Router, routes routes.Routes, env infrastructure.Env) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port 8080")
			go func() {
				routes.Setup()
				handler.Gin.Run(":8080")
			}()
			return nil

		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
