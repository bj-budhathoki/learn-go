package routes

import (
	"github.com/bj-budhathoki/learn-go/url_short/api/controllers"
	"github.com/bj-budhathoki/learn-go/url_short/infrastructure"
)

type TestRoutes struct {
	router        infrastructure.Router
	urlController controllers.UrlController
}

func NewTestRoutes(router infrastructure.Router, urlController controllers.UrlController) TestRoutes {
	return TestRoutes{
		router:        router,
		urlController: urlController,
	}
}

func (u TestRoutes) Setup() {
	urls := u.router.Gin.Group("/test")
	{
		urls.GET("", u.urlController.GetURL)
	}
}
