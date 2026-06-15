package handler

import (
	"strconv"

	"prophet/internal/logic"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderLogic *logic.OrderLogic
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{orderLogic: logic.NewOrderLogic()}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req struct {
		ProductID string `json:"product_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	userID, _ := c.Get("user_id")
	order, err := h.orderLogic.Create(userID.(uint), req.ProductID)
	if err != nil {
		jsonError(c, "商品不存在")
		return
	}

	jsonResponse(c, 0, order)
}

func (h *OrderHandler) List(c *gin.Context) {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	userID, _ := c.Get("user_id")
	orders, total, err := h.orderLogic.List(userID.(uint), status, page, limit)
	if err != nil {
		jsonError(c, "获取失败")
		return
	}

	jsonResponse(c, 0, gin.H{"orders": orders, "total": total})
}

func (h *OrderHandler) Confirm(c *gin.Context) {
	id := c.Param("id")
	if err := h.orderLogic.Confirm(id); err != nil {
		jsonError(c, "订单不存在")
		return
	}
	jsonResponse(c, 0, gin.H{"message": "确认成功"})
}
