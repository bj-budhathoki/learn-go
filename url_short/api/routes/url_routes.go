package routes

import (
	"github.com/bj-budhathoki/learn-go/url_short/api/controllers"
	"github.com/bj-budhathoki/learn-go/url_short/infrastructure"
)

type URLRoutes struct {
	router        infrastructure.Router
	urlController controllers.UrlController
}

func NewURLRoutes(router infrastructure.Router, urlController controllers.UrlController) URLRoutes {
	return URLRoutes{
		router:        router,
		urlController: urlController,
	}
}

func (u URLRoutes) Setup() {
	urls := u.router.Gin.Group("/urls")
	{
		urls.GET("", u.urlController.GetURL)
	}
}
