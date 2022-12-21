package handlers

import (
	"net/http"
	"pustaka-api/entity"

	"github.com/gin-gonic/gin"
)

func (h *bookHandler) GetAllBooksHandler(ctx *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []entity.BookResponse

	for _, book := range books {
		bookResponse := entity.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Subtitle:    book.Subtitle,
			Description: book.Description,
			Price:       book.Price,
			Rating:      book.Rating,
			Discount:    book.Discount,
		}

		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"error":  false,
		"data":   booksResponse,
	})
}
