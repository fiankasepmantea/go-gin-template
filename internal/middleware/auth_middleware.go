package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/fiankasepman/go-gin-template/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(401)
			return
		}

		token := strings.Replace(header, "Bearer ", "", 1)

		payload, err := auth.VerifyToken(token)

		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("user_id", payload.UserID)

		c.Next()
	}
}