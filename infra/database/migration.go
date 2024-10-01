package database

import "gin-example/models"

// Add list of model add for migrations
var migrationModels = []interface{}{
	&models.User{},
	&models.Product{},
	&models.Order{},
	&models.OrderItem{},
}
