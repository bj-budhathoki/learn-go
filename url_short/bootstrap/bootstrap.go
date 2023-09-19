package bootstrap

import (
	"context"
	"fmt"

	"github.com/bj-budhathoki/learn-go/url_short/api/controllers"
	"github.com/bj-budhathoki/learn-go/url_short/api/routes"
	"github.com/bj-budhathoki/learn-go/url_short/docs"
	"github.com/bj-budhathoki/learn-go/url_short/infrastructure"
	"go.uber.org/fx"
)

var Module = fx.Options(
	routes.Module,
	infrastructure.Module,
	controllers.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, env infrastructure.Env, handler infrastructure.Router, routes routes.Routes) {

	appStop := func(context.Context) error {
		fmt.Print("Stopping application")
		// database.DB.Disconnect(context.TODO())
		return nil
	}
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if env.Environment != "production" && env.HOST != "" {
					fmt.Println("Setting Swagger Host...")
					docs.SwaggerInfo.Host = env.HOST
				}
				routes.Setup()
				handler.Gin.Run(":8080")
			}()
			return nil
		},
		OnStop: appStop,
	})
}
