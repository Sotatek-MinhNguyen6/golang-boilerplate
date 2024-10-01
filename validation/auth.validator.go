package validation

import (
	"gin-example/dto" // Adjust the import path based on your project structure
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateLoginRequest validates the login request body
func ValidateLoginRequest(c *gin.Context) *dto.LoginRequest {
	var loginReq dto.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return &loginReq
}

func ValidateRegisterRequest(c *gin.Context) *dto.RegisterRequest {
	var registerReq dto.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return &registerReq
}

func ValidateRefreshTokenRequest(c *gin.Context) *string {
	var refreshReq dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&refreshReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	return &refreshReq.RefreshToken
}
