package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func QueryHandler2(ctx *gin.Context) {

	title := ctx.Query("title")
	subtitle := ctx.Query("subtitle")

	ctx.JSON(http.StatusOK, gin.H{
		"title":    title,
		"subtitle": subtitle,
	})
}

func BooksHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
func BooksHandler2(ctx *gin.Context) {

	id := ctx.Param("id")
	name := ctx.Param("name")

	ctx.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func WelcomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Welcome": "Hello you access nascript/go API 😁",
	})
}
func AboutHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"author": "Nasuiton",
		"bio":    "A Software Engineer",
	})
}

func QueryHandler(ctx *gin.Context) {

	title := ctx.Query("title")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

// POST

func AddNewBookHandler(ctx *gin.Context) {
	var bookInput book.BookInput
	err := ctx.ShouldBindJSON(&bookInput)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition: Field is %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":    bookInput.Title,
		"price":    bookInput.Price,
		"subtitle": bookInput.Subtitle,
	})
}
