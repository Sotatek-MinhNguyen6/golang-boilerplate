package routers

import (
	"gin-example/controllers"
	"gin-example/routers/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(route *gin.Engine) {
	orderController := controllers.OrderController{}
	orderRoute := route.Group("/orders")
	orderRoute.Use(middleware.JWTAuthMiddleware("user"))
	orderRoute.POST("", orderController.CreateOrder)
	orderRoute.POST("/:orderId/checkout", orderController.CheckoutOrder)
}
