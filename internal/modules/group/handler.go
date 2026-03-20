package group

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

func (h *Handler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		h.Error(c, err.Error())
		return
	}
	h.Success(c, data)
}

func (h *Handler) Create(c *gin.Context) {
	var req Group
	if !h.Bind(c, &req) {
		return
	}

	err := h.service.Create(&req)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, req)
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	var req Group
	if !h.Bind(c, &req) {
		return
	}

	req.GroupID = id

	err := h.service.Update(&req)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, req)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, gin.H{"deleted": id})
}