package group

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

	data, err := h.service.GetAll()
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, data)
}