package handlers

import (
	"net/http"
	"pustaka-api/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"pustaka-api/helpers"
)

func (h *bookHandler) GetBookByIdHandler(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.FindByID(id)

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
		"data":   booksResponse,
	})
}
