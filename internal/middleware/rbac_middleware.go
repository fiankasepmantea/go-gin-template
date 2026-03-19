package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RBACMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method

		var bypass int
		db.Table("endpoint").
			Select("bypass").
			Where("value = ? AND method = ?", path, method).
			Scan(&bypass)

		if bypass == 1 {
			c.Next()
			return
		}

		userID, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		var count int64
		db.Table("users u").
			Joins("JOIN groups g ON g.group_id = u.group_id").
			Joins("JOIN group_endpoint ge ON ge.group_id = g.group_id").
			Joins("JOIN endpoint e ON e.endpoint_id = ge.endpoint_id").
			Where("u.user_id = ? AND e.value = ? AND e.method = ?", userID, path, method).
			Count(&count)

		if count == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		c.Next()
	}
}