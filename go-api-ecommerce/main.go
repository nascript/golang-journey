package main

import (
	"log"
	"os"

	"github.com/nascript/golang-journey/controllers"
	"github.com/nascript/golang-journey/database"
	"github.com/nascript/golang-journey/middleware"
	"github.com/nascript/golang-journey/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("Port")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	// router.GET("/cartcheckout", app.BuyFromCart())

	log.Fatal(router.Run(":" + port))

}
