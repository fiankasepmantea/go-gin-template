package middleware

import (
	"net/http"
	"strings"

	"github.com/fiankasepman/go-gin-template/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		split := strings.Split(authHeader, " ")
		if len(split) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		token := split[1]

		payload, err := auth.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user_id", payload.UserID)
		c.Set("token_id", payload.TokenID)
		c.Set("domain_id", payload.DomainID)

		c.Next()
	}
}
