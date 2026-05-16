package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", 1)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
