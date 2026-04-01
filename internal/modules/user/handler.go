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
func (h *Handler) GetAll(c *gin.Context) {
	limit, _, offset := h.GetPagination(c)

	data, err := h.service.GetAll(limit, offset)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, data)
}
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
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, gin.H{"deleted": id})
}
func (h *Handler) Login(c *gin.Context) {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if !h.Bind(c, &req) {
		return
	}

	device := c.GetHeader("X-Device")
	ua := c.Request.UserAgent()
	ip := c.ClientIP()

	res, err := h.service.Login(req.Username, req.Password, device, ua, ip)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, res)
}
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

	h.Success(c, gin.H{"access_token": token})
}
func (h *Handler) Logout(c *gin.Context) {

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if !h.Bind(c, &req) {
		return
	}

	err := h.service.Logout(req.RefreshToken)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, "logout success")
}

func (h *Handler) LogoutAll(c *gin.Context) {

	userID := h.GetUserID(c)

	err := h.service.LogoutAll(userID)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, "logout all device success")
}
func (h *Handler) Devices(c *gin.Context) {

	userID := h.GetUserID(c)

	data, err := h.service.GetDevices(userID)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, data)
}
func (h *Handler) RevokeDevice(c *gin.Context) {

	userID := h.GetUserID(c)
	id := c.Param("id")

	err := h.service.RevokeDevice(userID, id)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"message": "device revoked",
	})
}