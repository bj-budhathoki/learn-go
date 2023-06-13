package routes

import (
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/api/controllers"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/infrastructure"
)

type AuthRoutes struct {
	router          infrastructure.Router
	userControllers controllers.UserControllers
}

func NewAuthRoutes(router infrastructure.Router,
	userControllers controllers.UserControllers) AuthRoutes {
	return AuthRoutes{
		router:          router,
		userControllers: userControllers,
	}
}

// Setup activityParticipation routes
func (c AuthRoutes) Setup() {
	router := c.router.Gin.Group("/auth")
	{
		router.POST("/login", c.userControllers.CreateNewUser)
		router.POST("/register", c.userControllers.CreateNewUser)
	}
}
