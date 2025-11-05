package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtManager *JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		authTokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		authTokenStr = strings.TrimSpace(authTokenStr)

		authToken, err := jwtManager.Parse(authTokenStr)
		if err != nil || !authToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()

			return
		}

		if claims, ok := authToken.Claims.(jwt.MapClaims); ok {
			userID := int(claims["user_id"].(float64))
			c.Set("user_id", userID)
		}

		c.Next()
	}
}
