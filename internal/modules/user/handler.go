package user

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

// GET ALL
func (h *Handler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		h.Error(c, err.Error())
		return
	}
	h.Success(c, data)
}

// CREATE
func (h *Handler) Create(c *gin.Context) {
	var req User
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

// UPDATE
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	var req User
	if !h.Bind(c, &req) {
		return
	}

	req.UserID = id

	err := h.service.Update(&req)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, req)
}

// DELETE
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, gin.H{"deleted": id})
}

// LOGIN
func (h *Handler) Login(c *gin.Context) {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
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

// ME
func (h *Handler) Me(c *gin.Context) {

	userID := h.GetUserID(c) // pakai BaseHandler

	h.Success(c, gin.H{
		"user_id": userID,
	})
}

func (h *Handler) Refresh(c *gin.Context) {

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if !h.Bind(c, &req) {
		return
	}

	token, err := h.service.Refresh(req.RefreshToken)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"access_token": token,
	})
}
func (h *Handler) Logout(c *gin.Context) {

	userID := h.GetUserID(c)

	err := h.service.Logout(userID)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"message": "logout success",
	})
}