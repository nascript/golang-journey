package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nascript/golang-journey/controllers"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/users/signup", controllers.SignUp())
	incommingRoutes.POST("/users/login", controllers.Login())
	incommingRoutes.POST("/admin/addproducts", controllers.ProductViewerAdmin())
	incommingRoutes.GET("/users/productview", controllers.SearchProduct())
	incommingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
