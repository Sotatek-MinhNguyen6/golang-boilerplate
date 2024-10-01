package routers

import (
	"gin-example/controllers"
	"gin-example/routers/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(route *gin.Engine) {
	controller := controllers.ProductController{}
	productRoute := route.Group("/products")
	productRoute.Use(middleware.JWTAuthMiddleware()) // Apply JWT middleware
	productRoute.POST("", controller.CreateProduct)
	productRoute.GET("", controller.ListProducts)
}
