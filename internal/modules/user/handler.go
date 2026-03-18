package user

import (
	"github.com/fiankasepman/go-gin-template/internal/base"
	"github.com/fiankasepman/go-gin-template/internal/middleware"
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

func (h *Handler) Login(c *gin.Context) {

	var req struct {
		Username string
		Password string
	}

	if !h.Bind(c, &req) {
		return
	}

	res, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, res)
}

func (h *Handler) Me(c *gin.Context) {

	userID := middleware.GetUserID(c)

	h.Success(c, gin.H{
		"user_id": userID,
	})
}