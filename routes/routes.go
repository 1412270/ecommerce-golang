package routes

import (
	"github.com/1412270/ecommerce-golang/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("admin/addproduct", controllers.ProductViewAdmin())
	incomingRoutes.GET("users/productview", controllers.SearchProduct())
	incomingRoutes.GET("users/search", controllers.SearchProductByQuery())
}
