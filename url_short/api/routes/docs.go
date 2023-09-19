package routes

import (
	"github.com/bj-budhathoki/learn-go/url_short/infrastructure"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type DocsRoutes struct {
	router infrastructure.Router
	env    infrastructure.Env
}

func NewDocsRoute(router infrastructure.Router, env infrastructure.Env) DocsRoutes {
	return DocsRoutes{
		router: router,
		env:    env,
	}
}

func (d DocsRoutes) Setup() {
	if d.env.Environment != "production" {
		swagger := d.router.Gin.Group("/docs")
		{
			swagger.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		}
	}
}
