package routers

import (
	"gin-example/routers/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routes() *gin.Engine {
	environment := viper.GetBool("DEBUG")
	if environment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.HttpExceptionHandlerMiddleware())
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})
	// existing product routes setup
	ProductRoutes(router)
	AuthRoutes(router)
	UserRoute(router)
	OrderRoutes(router) // Add this line to include order routes

	return router
}
