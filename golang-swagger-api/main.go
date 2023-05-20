package main

import (
	"fmt"

	"github.com/bj-budhathoki/learn-go/golang-swagger-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	routes.RegisterBookRoutes(r)
	r.Run()
}
