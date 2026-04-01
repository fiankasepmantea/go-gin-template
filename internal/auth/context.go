package auth

import "github.com/gin-gonic/gin"
func GetDomainID(c *gin.Context) int {
	val, ok := c.Get("domain_id")
	if !ok {
		return 0
	}
	return val.(int)
}