package middleware

import (
	"errors"
	"net/http"

	"github.com/fiankasepman/go-gin-template/configs"
	"github.com/fiankasepman/go-gin-template/internal/cache"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RBACMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.FullPath()
		method := c.Request.Method

		// ================== USER CONTEXT ==================
		userID := c.GetString("user_id")
		domainID := c.GetInt("domain_id")

		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

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

		// ================== BYPASS ==================
		if endpoint.Bypass == 1 {
			c.Next()
			return
		}

		// ================== SUPER ADMIN ==================
		var isAdmin int

		err = db.Table("users").
			Select("is_admin").
			Where("user_id = ? AND domain_id = ?", userID, domainID).
			Scan(&isAdmin).Error

		if err == nil && isAdmin == 1 {
			c.Next()
			return
		}

		// ================== REDIS CACHE ==================
		key := "rbac:" + userID + ":" + endpoint.EndpointID

		val, err := cache.RDB.Get(cache.Ctx, key).Result()

		if err == nil {
			// cache hit
			if val == "1" {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		// kalau bukan karena key tidak ada → redis error
		if err != nil && !errors.Is(err, redis.Nil) {
			// log optional
			// log.Println("redis error:", err)
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

		_ = cache.RDB.Set(cache.Ctx, key, cacheValue, configs.AccessTokenDuration).Err()

		// ================== FINAL CHECK ==================
		if count == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		c.Next()
	}
}