package main

import (
	"gin-example/config"
	"gin-example/infra/database"
	"gin-example/infra/logger"
	"gin-example/routers"
	"github.com/spf13/viper"
	"time"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Dhaka")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	masterDSN := config.DbConfiguration()
	if err := database.DBConnection(masterDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}
	router := routers.Routes()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))
}
