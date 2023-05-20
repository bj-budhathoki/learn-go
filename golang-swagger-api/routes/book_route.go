package routes

import (
	"github.com/bj-budhathoki/learn-go/golang-swagger-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine) {
	routes := router.Group("/books")
	routes.GET("/", controllers.GetBooks)
	routes.GET("/:id", controllers.GetBook)
}
