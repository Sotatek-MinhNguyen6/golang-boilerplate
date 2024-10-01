package controllers

import (
	"gin-example/infra/database"
	"gin-example/infra/logger"
	"gin-example/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (ctrl *ProductController) CreateProduct(ctx *gin.Context) {
	newProduct := new(models.Product)
	err := ctx.BindJSON(&newProduct)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&newProduct).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &newProduct)
}

func (ctrl *ProductController) ListProducts(ctx *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	ctx.JSON(http.StatusOK, gin.H{"data": products})
}
