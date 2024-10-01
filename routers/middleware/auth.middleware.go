package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// JWTAuthMiddleware checks for a valid JWT and ensures the user has the correct role if specified.
func JWTAuthMiddleware(targetRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwtv4.Parse(tokenString, func(token *jwtv4.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwtv4.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwtv4.MapClaims); ok && token.Valid {
			if len(targetRoles) > 0 {
				role := claims["role"].(string)
				validRole := false
				for _, targetRole := range targetRoles {
					if role == targetRole {
						validRole = true
						break
					}
				}
				if !validRole {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
					c.Abort()
					return
				}
			}
			c.Set("userId", claims["userId"])
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
