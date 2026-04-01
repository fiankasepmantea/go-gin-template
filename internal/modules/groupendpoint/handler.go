package groupendpoint

import (
	"github.com/fiankasepman/go-gin-template/internal/base"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	base.BaseHandler
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Assign(c *gin.Context) {

	var req struct {
		GroupID    string `json:"group_id"`
		EndpointID string `json:"endpoint_id"`
	}

	if !h.Bind(c, &req) {
		return
	}

	err := h.service.Assign(req.GroupID, req.EndpointID)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, "assigned")
}

func (h *Handler) Remove(c *gin.Context) {

	var req struct {
		GroupID    string `json:"group_id"`
		EndpointID string `json:"endpoint_id"`
	}

	if !h.Bind(c, &req) {
		return
	}

	err := h.service.Remove(req.GroupID, req.EndpointID)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, "removed")
}

func (h *Handler) GetByGroup(c *gin.Context) {

	groupID := c.Param("group_id")

	data, err := h.service.GetByGroup(groupID)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, data)
}