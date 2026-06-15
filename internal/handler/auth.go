package handler

import (
	"net/http"

	"prophet/internal/logic"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authLogic *logic.AuthLogic
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{authLogic: logic.NewAuthLogic()}
}

func (h *AuthHandler) AnonymousInit(c *gin.Context) {
	var req struct {
		DeviceID string `json:"device_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	user, token, err := h.authLogic.AnonymousInit(req.DeviceID)
	if err != nil {
		jsonError(c, "服务器错误")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get("user_id")
	user, err := h.authLogic.GetMe(userID)
	if err != nil {
		jsonError(c, "用户不存在")
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"user": user}})
}
