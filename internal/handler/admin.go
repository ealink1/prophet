package handler

import (
	"prophet/internal/logic"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminLogic *logic.AdminLogic
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{adminLogic: logic.NewAdminLogic()}
}

func (h *AdminHandler) Dashboard(c *gin.Context) {
	data, err := h.adminLogic.Dashboard()
	if err != nil {
		jsonError(c, "获取失败")
		return
	}
	jsonResponse(c, 0, data)
}
