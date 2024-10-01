package validation

import (
	"gin-example/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateCreateOrderRequest(c *gin.Context) *dto.CreateOrderRequest {
	var createOrderReq dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&createOrderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	return &createOrderReq
}
