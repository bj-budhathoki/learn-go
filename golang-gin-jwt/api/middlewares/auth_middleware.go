package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return AuthMiddleware{}
}
func (cc AuthMiddleware) Authenticate() gin.HandlerFunc {
	fmt.Println("Im a dummy!")
	return func(c *gin.Context) {
		c.Next()
	}
}
