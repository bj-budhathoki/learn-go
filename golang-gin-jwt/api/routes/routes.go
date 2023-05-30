package routes

import (
	"log"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoutes),
)

type Route interface {
	Setup()
}
type Routes []Route

func NewRoutes(userRoutes UserRoutes) Routes {
	return Routes{
		userRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	log.Print("routes")
	for _, route := range r {
		route.Setup()
	}
}
