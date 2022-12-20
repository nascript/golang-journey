package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", welcomeHandler)
	router.GET("/about", aboutHandler) // get data
	router.GET("/books/:id", booksHandler) //get by single id
	router.GET("/books/:id/:name", booksHandler2) //get by multiple id
	// router.GET("/query", queryHandler) // get by single query
	router.GET("/query", queryHandler2) // get by multiple query

	
	router.Run(":7777")
}


// func queryHandler(ctx *gin.Context) {

// 	title := ctx.Query("title")

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"title" : title,
// 	})
// }

func queryHandler2(ctx *gin.Context) {

	title := ctx.Query("title")
	subtitle := ctx.Query("subtitle")

	ctx.JSON(http.StatusOK, gin.H{
		"title" : title,
		"subtitle" : subtitle,
	})
}

func booksHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id" : id,
	})
}
func booksHandler2(ctx *gin.Context) {

	id := ctx.Param("id")
	name := ctx.Param("name")

	ctx.JSON(http.StatusOK, gin.H{
		"id" : id,
		"name" : name,
	})
}


func welcomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
			"Welcome" : "Hello you access nascript/go API 😁",
		})
}
func aboutHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
			"author" : "Nasuiton",
			"bio" : "A Software Engineer",
		})
}
