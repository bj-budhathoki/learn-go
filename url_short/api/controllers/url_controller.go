package controllers

import (
	"github.com/bj-budhathoki/learn-go/url_short/infrastructure"
	"github.com/gin-gonic/gin"
)

type UrlController struct {
	env infrastructure.Env
}

func NewUrlController(env infrastructure.Env) UrlController {
	return UrlController{
		env: env,
	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (cc UrlController) GetURL(c *gin.Context) {
	c.JSON(200, gin.H{
		"url": "example.com",
	})
}
