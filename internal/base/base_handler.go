package base

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct{}

func (h *BaseHandler) JSON(c *gin.Context, httpCode int, success bool, message string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"success": success,
		"message": message,
		"data":    data,
	})
}

func (h *BaseHandler) Success(c *gin.Context, data interface{}) {
	h.JSON(c, http.StatusOK, true, "success", data)
}

func (h *BaseHandler) Created(c *gin.Context, data interface{}) {
	h.JSON(c, http.StatusCreated, true, "created", data)
}

func (h *BaseHandler) Error(c *gin.Context, message string) {
	h.JSON(c, http.StatusBadRequest, false, message, nil)
}

func (h *BaseHandler) NotFound(c *gin.Context, message string) {
	h.JSON(c, http.StatusNotFound, false, message, nil)
}

func (h *BaseHandler) Unauthorized(c *gin.Context, message string) {
	h.JSON(c, http.StatusUnauthorized, false, message, nil)
}

func (h *BaseHandler) Bind(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		h.Error(c, err.Error())
		return false
	}
	return true
}

func (h *BaseHandler) GetUserID(c *gin.Context) string {
	userID, exists := c.Get("user_id")
	if !exists || userID == nil {
		return ""
	}
	val, ok := userID.(string)
	if !ok {
		return ""
	}
	return val
}

func (h *BaseHandler) GetPagination(c *gin.Context) (int, int, int) {
	limit := 10
	page := 1

	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}
	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	return limit, page, offset
}
func (h *BaseHandler) GetTokenID(c *gin.Context) string {
	val, exists := c.Get("token_id")
	if !exists || val == nil {
		return ""
	}

	str, ok := val.(string)
	if !ok {
		return ""
	}

	return str
}