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

		// ================== CHECK ENDPOINT ==================
		var endpoint struct {
			EndpointID string
			Bypass     int
		}

		err := db.Table("endpoint").
			Select("endpoint_id, bypass").
			Where("value = ? AND method = ?", path, method).
			Take(&endpoint).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "endpoint not registered"})
			return
		}

		// bypass endpoint
		if endpoint.Bypass == 1 {
			c.Next()
			return
		}

		// ================== GET USER CONTEXT ==================
		userIDVal, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		userID := userIDVal.(string)

		domainIDVal, _ := c.Get("domain_id")
		domainID := domainIDVal.(int)

		// ================== CHECK ADMIN ==================
		var isAdmin int
		err = db.Table("users").
			Select("is_admin").
			Where("user_id = ? AND domain_id = ?", userID, domainID).
			Scan(&isAdmin).Error

		if err == nil && isAdmin == 1 {
			c.Next()
			return
		}

		// ================== CACHE CHECK ==================
		key := BuildKey(userID, endpoint.EndpointID)

		if allowed, ok := RBACCache[key]; ok {
			if allowed {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		// ================== CHECK PERMISSION ==================
		var count int64

		err = db.Table("users u").
			Joins("JOIN groups g ON g.group_id = u.group_id").
			Joins("JOIN group_endpoint ge ON ge.group_id = g.group_id").
			Where("u.user_id = ? AND u.domain_id = ? AND ge.endpoint_id = ?", userID, domainID, endpoint.EndpointID).
			Count(&count).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "rbac error"})
			return
		}
		
		RBACCache[key] = count > 0
		
		if count == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		
		c.Next()
	}
}
