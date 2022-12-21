package handlers

import (
	"fmt"
	"net/http"
	"pustaka-api/entity"
	"pustaka-api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *bookHandler) AddNewBookHandler(ctx *gin.Context) {
	var bookRequest entity.BookRequest
	err := ctx.ShouldBindJSON(&bookRequest)

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
	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var booksResponse []entity.BookResponse

	bookResponse := helpers.ConvertToBookResponse(book)

	booksResponse = append(booksResponse, bookResponse)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"error":  false,
		"data":   bookResponse,
	})
}
