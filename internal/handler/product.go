package handler

import (
	"prophet/internal/logic"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) List(c *gin.Context) {
	products, err := logic.ListActiveProducts()
	if err != nil {
		jsonError(c, "获取失败")
		return
	}
	jsonResponse(c, 0, products)
}
