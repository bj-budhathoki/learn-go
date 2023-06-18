package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/bj-budhathoki/learn-go/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	if _, err := os.Stat("public/single"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/single", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	if _, err := os.Stat("public/multiple"); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("public/multiple", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
func main() {
	router := gin.Default()

	router.GET("/health_checker", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "success", "message": "How to Upload Single and Multiple Files in Golang"})
	})
	router.POST("/upload/single", controllers.UploadFile)
	router.POST("/upload/multiple", controllers.UploadFiles)
	router.StaticFS("/images", http.Dir("public"))
	router.Run(":8000")
}
