package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router -> Gin Router
type Router struct {
	Gin *gin.Engine
}

func NewRouter(env Env) Router {
	httpRouter := gin.Default()
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Project Easyy 📺 API Up and Running"})
	})
	return Router{
		Gin: httpRouter,
	}
}
