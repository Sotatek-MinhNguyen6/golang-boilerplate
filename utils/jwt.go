package utils

import (
	"errors"
	"gin-example/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var mySigningKey = []byte(viper.GetString("JWT_SECRET")) // Use a secure method to store the secret key

// GenerateJWT generates an access token and a refresh token
func GenerateJWT(user *models.User) (string, string, error) {
	// Access token
	accessTokenClaims := jwt.MapClaims{
		"authorized": true,
		"userId":     user.ID,
		"username":   user.Username,
		"role":       user.Role,
		"exp":        time.Now().Add(time.Minute * time.Duration(viper.GetInt("JWT_ACCESS_TOKEN_EXPIRATION"))).Unix(), // Access token expires in 15 minutes
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	// Refresh token
	refreshTokenClaims := jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * time.Duration(viper.GetInt("JWT_REFRESH_TOKEN_EXPIRATION"))).Unix(), // Refresh token expires in 7 days
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func VerifyJWT(tokenString string) (id uint, role string, err error) {
	// Define a struct to hold the expected claims
	type Claims struct {
		UserID uint `json:"userId"`
		Role   string
		jwt.RegisteredClaims
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return 0, "", err
	}

	// Validate the token and ensure it is not expired
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, claims.Role, nil
	}

	return 0, "", errors.New("invalid token")
}
