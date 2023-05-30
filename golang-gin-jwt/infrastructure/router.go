package infrastructure

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin *gin.Engine
}

func NewRouter(env Env) Router {
	log.Print("route handler")
	httpRouter := gin.Default()
	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "API Health Check: Up and Running"})
	})
	return Router{
		Gin: httpRouter,
	}
}
