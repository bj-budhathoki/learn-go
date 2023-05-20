package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: 1, Title: "HTML and CSS: Design and Build Websites", Author: "Jon Duckett", Quantity: 2},
	{ID: 2, Title: "Coding All-In-One for Dummies", Author: " DK", Quantity: 5},
}

func GetBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data":    books,
		"message": "success",
	})
}

// create new book
func CreateBook(ctx *gin.Context) {
	var newBook Book
	if err := ctx.BindJSON(&newBook); err != nil {
		ctx.AbortWithError(500, err)
	}
	books = append(books, newBook)
	ctx.IndentedJSON(http.StatusCreated, newBook)
}

// get single book by id
func getBookById(id int) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func GetBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book, err := getBookById(id)
	if err != nil {
		ctx.AbortWithStatus(500)
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"data":    book,
		"message": "success",
	})
}

// delete book

func DeleteBook(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
