package base

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct{}

func (h *BaseHandler) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *BaseHandler) Error(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"message": message,
	})
}

func (h *BaseHandler) GetUserID(c *gin.Context) string {
	userID, _ := c.Get("user_id")
	if userID == nil {
		return ""
	}
	return userID.(string)
}

func (h *BaseHandler) Bind(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		h.Error(c, err.Error())
		return false
	}
	return true
}

func (h *BaseHandler) GetPagination(c *gin.Context) (int, int) {
	limit := 10
	page := 1

	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}

	offset := (page - 1) * limit
	return limit, offset
}