package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllers struct{}

func NewUserControllers() UserControllers {
	return UserControllers{}
}
func (c UserControllers) CreateNewUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})

}
