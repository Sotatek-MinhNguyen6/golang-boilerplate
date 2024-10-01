package routers

import (
	"gin-example/controllers"
	"gin-example/routers/middleware"

	"github.com/gin-gonic/gin"
)

// Setup routes
func UserRoute(route *gin.Engine) {
	controller := controllers.UserController{}
	adminRoutes := route.Group("/admin")
	adminRoutes.Use(middleware.JWTAuthMiddleware("admin")) // Use the AdminOnly middleware from the new file
	{
		adminRoutes.GET("/users", controller.ListUsers) // Changed 'listUser' to 'listUsers'
	}
}
