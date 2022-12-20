package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

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

	db.AutoMigrate(&book.Book{})

	// CRUD

	// book := book.Book{} // for single data
	var books []book.Book
	// CREATE
	// book.Title = "Cara mencintaimu"
	// book.Description = "Lorem ipsum Dollor Lorem ipsum Dollor Lorem ipsum Dollor"
	// book.Price = 10302
	// book.Discount = 20
	// book.Rating = 2

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("====================")
	// 	fmt.Println("Error Add Data Book")
	// 	fmt.Println("====================")
	// }

	// GET  SINGLE DATA
	// err = db.Debug().First(&book).Error
	// if err != nil {
	// 	fmt.Println("===== SINGLE DATA =====")
	// 	fmt.Println("Error Finding Data Book")
	// 	fmt.Println("====================")
	// }

	// fmt.Println("====================")
	// fmt.Println("Title :", book.Title)
	// fmt.Println("book Object %v", book)
	// fmt.Println("====================")

	// GET ALL DATA
	err = db.Debug().Find(&books).Error
	if err != nil {
		fmt.Println("====================")
		fmt.Println("Error Finding Data Book")
		fmt.Println("====================")
	}

	fmt.Println("=====ALL DATA =====")
	for _, book := range books {
		fmt.Println("Title :", book.Title)
		fmt.Println("book Object %v", book)
		fmt.Println("====================")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	// GET METHOD =======================
	v1.GET("/", handler.WelcomeHandler)
	v1.GET("/about", handler.AboutHandler)            // get data
	v1.GET("/books/:id", handler.BooksHandler)        //get by single id
	v1.GET("/books/:id/:name", handler.BooksHandler2) //get by multiple id
	// v1.GET("/query", queryHandler) // get by single query
	v1.GET("/query", handler.QueryHandler2) // get by multiple query

	// POST METHOD
	v1.POST("/book", handler.AddNewBookHandler)

	router.Run(":7777")
}
