package controllers

import (
	"gin-example/infra/database"
	"gin-example/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (c UserController) ListUsers(ctx *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
