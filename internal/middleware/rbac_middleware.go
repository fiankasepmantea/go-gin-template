package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fiankasepman/go-gin-template/internal/cache"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RBACMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.FullPath()
		method := c.Request.Method

		userID := c.GetString("user_id")
		domainID := c.GetInt("domain_id")

		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// ================== GET ENDPOINT ==================
		bypass, endpointID, err := getEndpoint(db, path, method)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "endpoint not registered"})
			return
		}

		if bypass {
			c.Next()
			return
		}

		// ================== SUPER ADMIN ==================
		if isSuperAdmin(db, userID, domainID) {
			c.Next()
			return
		}

		// ================== REDIS ==================
		key := fmt.Sprintf("rbac:%d:%s:%s", domainID, userID, endpointID)

		val, err := cache.RDB.Get(cache.Ctx, key).Result()
		if err == nil {
			if val == "1" {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		// ================== DB CHECK ==================
		allowed, err := checkPermission(db, userID, domainID, endpointID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "rbac error"})
			return
		}

		// ================== CACHE ==================
		cacheValue := "0"
		if allowed {
			cacheValue = "1"
		}

		cache.RDB.Set(cache.Ctx, key, cacheValue, 10*time.Minute)

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}

		c.Next()
	}
}

func getEndpoint(db *gorm.DB, path, method string) (bool, string, error) {

	var endpoint struct {
		EndpointID string
		Bypass     int
	}

	err := db.Table("endpoint").
		Select("endpoint_id, bypass").
		Where("value = ? AND method = ?", path, method).
		Take(&endpoint).Error

	if err != nil {
		return false, "", err
	}

	return endpoint.Bypass == 1, endpoint.EndpointID, nil
}

func isSuperAdmin(db *gorm.DB, userID string, domainID int) bool {
	var isAdmin int

	db.Table("users").
		Select("is_admin").
		Where("user_id = ? AND domain_id = ?", userID, domainID).
		Scan(&isAdmin)

	return isAdmin == 1
}

func checkPermission(db *gorm.DB, userID string, domainID int, endpointID string) (bool, error) {

	var count int64

	err := db.Table("users u").
		Joins("JOIN groups g ON g.group_id = u.group_id AND g.domain_id = u.domain_id").
		Joins("JOIN group_endpoint ge ON ge.group_id = g.group_id").
		Where(`
			u.user_id = ? 
			AND u.domain_id = ? 
			AND ge.endpoint_id = ?
		`, userID, domainID, endpointID).
		Count(&count).Error

	return count > 0, err
}