package controllers

import (
	"net/http"

	"github.com/bj-budhathoki/learn-go/golang-swagger-api/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) {
	books := []models.BookModel{
		{ID: 1, Title: "HTML and CSS: Design and Build Websites", Author: "Jon Duckett", Quantity: 2},
		{ID: 2, Title: "Coding All-In-One for Dummies", Author: " DK", Quantity: 5},
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data":    books,
		"message": "success",
	})
}
func GetBook(ctx *gin.Context) {
	book := models.BookModel{
		ID: 1, Title: "HTML and CSS: Design and Build Websites", Author: "Jon Duckett", Quantity: 2,
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data":    book,
		"message": "success",
	})
}
