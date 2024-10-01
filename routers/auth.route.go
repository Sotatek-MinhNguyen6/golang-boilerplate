package routers

import (
	"gin-example/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	authController := controllers.AuthController{}
	authGroup := route.Group("/auth")
	{
		authGroup.POST("/register", authController.RegisterUser)
		authGroup.POST("/login", authController.LoginUser)
		authGroup.POST("/refresh", authController.RefreshToken)
	}
}
