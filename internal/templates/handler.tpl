package {{.Package}}

import (
	"github.com/gin-gonic/gin"
	"github.com/fiankasepman/go-gin-template/internal/base"
)

type Handler struct {
	base.BaseHandler
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(c *gin.Context) {
	h.Success(c, "{{.Package}} list")
}