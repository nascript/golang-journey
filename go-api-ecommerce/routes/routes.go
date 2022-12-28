package routes

import (
	"github.com/nascript/golang-journey/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/users/signup", controllers.Signup())
	incommingRoutes.POST("/users/login", controllers.Login())
	incommingRoutes.POST("/admin/addproducts", controllers.ProductViewerAdmin())
	incommingRoutes.GET("/users/productview", controllers.SearchProduct())
	incommingRoutes.GET("/users/search", controllers.SearchProductProductByQuery())
}