package handler

import (
	"time"

	"prophet/internal/logic"

	"github.com/gin-gonic/gin"
)

type AlmanacHandler struct{}

func NewAlmanacHandler() *AlmanacHandler {
	return &AlmanacHandler{}
}

func (h *AlmanacHandler) Today(c *gin.Context) {
	var req struct {
		Date string `json:"date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Date == "" {
		req.Date = time.Now().Format("2006-01-02")
	}

	data := logic.ComputeAlmanac(req.Date)
	jsonResponse(c, 0, data)
}

func (h *AlmanacHandler) Week(c *gin.Context) {
	data := logic.ComputeWeekAlmanac()
	jsonResponse(c, 0, gin.H{"items": data})
}
