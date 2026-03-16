package user

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

func (h *Handler) Login(c *gin.Context) {

	var req struct {
		Username string
		Password string
	}

	if !h.Bind(c, &req) {
		return
	}

	user, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		h.Error(c, "invalid credential")
		return
	}

	h.Success(c, user)
}