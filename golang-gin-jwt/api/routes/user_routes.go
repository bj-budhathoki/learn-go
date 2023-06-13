package routes

import (
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/api/controllers"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/api/middlewares"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/infrastructure"
)

type UserRoutes struct {
	router          infrastructure.Router
	userControllers controllers.UserControllers
	middlewares     middlewares.AuthMiddleware
}

func NewUserRoutes(router infrastructure.Router,
	userControllers controllers.UserControllers, middlewares middlewares.AuthMiddleware) UserRoutes {
	return UserRoutes{
		router:          router,
		userControllers: userControllers,
		middlewares:     middlewares,
	}
}

// Setup activityParticipation routes
func (c UserRoutes) Setup() {
	router := c.router.Gin.Group("/users").Use(c.middlewares.Authenticate())
	{
		router.GET("/", c.userControllers.CreateNewUser)
	}
}
