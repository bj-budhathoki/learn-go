package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err :%s", err.Error()))
	}
	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "http://localhost:8000/images/single/" + filename
	out, err := os.Create("public/single/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
}

func UploadFiles(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["images"]
	filePaths := []string{}
	for _, file := range files {
		fileExt := filepath.Ext(file.Filename)
		originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
		now := time.Now()
		filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
		filePath := "http://localhost:8000/images/multiple/" + filename

		filePaths = append(filePaths, filePath)
		out, err := os.Create("./public/multiple/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		readerFile, _ := file.Open()
		_, err = io.Copy(out, readerFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"filepath": filePaths})
}
