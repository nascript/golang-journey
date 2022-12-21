package main

import (
	"fmt"
	"log"
	"pustaka-api/entity"
	"pustaka-api/handlers"
	"pustaka-api/repositories"
	"pustaka-api/services"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection Error")
	}
	fmt.Println("DB Connected")

	db.AutoMigrate(&entity.Book{})

	bookRepository := repositories.NewRepository(db)
	bookService := services.NewService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	// GET METHOD =======================
	v1.GET("/", bookHandler.WelcomeHandler)
	// v1.GET("/about", bookHandler.AboutHandler)            // get data
	// v1.GET("/books/:id", bookHandler.BooksHandler)        //get by single id
	// v1.GET("/books/:id/:name", bookHandler.BooksHandler2) //get by multiple id
	// v1.GET("/query", queryHandler) // get by single query
	// v1.GET("/query", bookHandler.QueryHandler2) // get by multiple query

	v1.POST("/book", bookHandler.AddNewBookHandler)
	v1.GET("/books", bookHandler.GetAllBooksHandler)
	v1.GET("/book/:id", bookHandler.GetBookByIdHandler)

	router.Run(":7777")
}
