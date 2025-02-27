package middleware

import (
	"golang-rnd/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware checks the Authorization header
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		claims, err := controllers.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store claims in context for reuse in controllers
		c.Set("username", claims.Username)
		c.Next()
	}
}
