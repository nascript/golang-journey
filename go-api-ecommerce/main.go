package main


import (
	"github.com/nascript/golang-journey/ecommerce/controllers"
	"github.com/nascript/golang-journey/ecommerce/routes"
	"github.com/nascript/golang-journey/ecommerce/middleware"
	"github.com/nascript/golang-journey/ecommerce/database"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("Port")

	if port == "" {
		port == "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":", + port))

}

