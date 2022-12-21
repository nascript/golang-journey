package handlers

import (
	"net/http"
	"pustaka-api/entity"
	"pustaka-api/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *bookHandler) DeleteBookHandler(ctx *gin.Context) {

	// GET ID
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Delete(int(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var booksResponse []entity.BookResponse

	bookResponse := helpers.ConvertToBookResponse(book)

	booksResponse = append(booksResponse, bookResponse)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"error":   false,
		"data":    bookResponse,
		"message": "Delete book successfully",
	})
}
