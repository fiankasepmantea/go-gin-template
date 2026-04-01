package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/fiankasepman/go-gin-template/internal/cache"
	"gorm.io/gorm"
)

func RBACMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.FullPath()
		method := c.Request.Method

		// ================== GET ENDPOINT ==================
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

		if endpoint.Bypass == 1 {
			c.Next()
			return
		}

		// ================== USER CONTEXT ==================
		userID := c.GetString("user_id")
		domainID := c.GetInt("domain_id")

		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// ================== ADMIN CHECK ==================
		var isAdmin int
		db.Table("users").
			Select("is_admin").
			Where("user_id = ? AND domain_id = ?", userID, domainID).
			Scan(&isAdmin)

		if isAdmin == 1 {
			c.Next()
			return
		}

		// ================== REDIS CACHE ==================
		key := "rbac:" + userID + ":" + endpoint.EndpointID

		val, err := cache.RDB.Get(cache.Ctx, key).Result()

		if err == nil {
			if val == "1" {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		// ================== DB FALLBACK ==================
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

		// ================== SET CACHE ==================
		cacheValue := "0"
		if count > 0 {
			cacheValue = "1"
		}

		cache.RDB.Set(cache.Ctx, key, cacheValue, 10*time.Minute)

		if count == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		c.Next()
	}
}