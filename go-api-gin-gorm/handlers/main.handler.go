package handlers

import (
	"net/http"
	"pustaka-api/services"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) WelcomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Welcome": "Hello you access nascript/go API 😁",
	})
}
