package controllers

import (
	"gin-example/infra/database"
	"gin-example/models"
	"gin-example/utils" // Ensure utils package is correctly imported
	"gin-example/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func (ctrl *AuthController) RegisterUser(ctx *gin.Context) {
	body := validation.ValidateRegisterRequest(ctx)
	if body == nil {
		return
	}

	var isUserExisted bool
	err := database.DB.Model(&models.User{}).Select("count(id) > 0").Where("username = ?", body.Username).Find(&isUserExisted).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isUserExisted {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User already existed"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{}
	user.Username = body.Username
	user.Role = body.Role
	user.Password = string(hashedPassword)
	if err := database.DB.Model(&models.User{}).Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (ctrl *AuthController) LoginUser(ctx *gin.Context) {
	body := validation.ValidateLoginRequest(ctx)
	if body == nil {
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", body.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateJWT(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Login successful",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (ctrl *AuthController) RefreshToken(ctx *gin.Context) {
	refreshToken := validation.ValidateRefreshTokenRequest(ctx)
	if refreshToken == nil {
		return
	}

	id, _, err := utils.VerifyJWT(*refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	user := models.User{}
	if err := database.DB.Model(&models.User{}).First(&user, id).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	newAccessToken, newRefreshToken, err := utils.GenerateJWT(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	})
}
