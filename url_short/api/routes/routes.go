package routes

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewURLRoutes),
	fx.Provide(NewDocsRoute),
	fx.Provide(NewTestRoutes),
)

// Route interface
type Route interface {
	Setup()
}

// Routes contains multiple routes
type Routes []Route

func NewRoutes(urlRoutes URLRoutes, docsRoutes DocsRoutes, testRoutes TestRoutes) Routes {
	return Routes{
		urlRoutes,
		docsRoutes,
		testRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
