package routes

import (
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/api/controllers"
	"github.com/bj-budhathoki/learn-go/golang-gin-jwt/infrastructure"
)

type UserRoutes struct {
	router          infrastructure.Router
	userControllers controllers.UserControllers
}

func NewUserRoutes(router infrastructure.Router,
	userControllers controllers.UserControllers) UserRoutes {
	return UserRoutes{
		router:          router,
		userControllers: userControllers,
	}
}

// Setup activityParticipation routes
func (c UserRoutes) Setup() {
	router := c.router.Gin.Group("/users")
	{
		router.POST("/", c.userControllers.CreateNewUser)
	}
}
