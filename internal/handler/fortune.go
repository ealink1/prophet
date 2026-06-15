package handler

import (
	"prophet/internal/logic"

	"github.com/gin-gonic/gin"
)

type FortuneHandler struct{}

func NewFortuneHandler() *FortuneHandler {
	return &FortuneHandler{}
}

func (h *FortuneHandler) LotteryDraw(c *gin.Context) {
	var req struct {
		Master   string `json:"master"`
		Question string `json:"question"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	result := logic.DrawLottery(req.Master)
	jsonResponse(c, 0, result)
}

func (h *FortuneHandler) BaziAnalyze(c *gin.Context) {
	var req struct {
		Year    int    `json:"year"`
		Month   int    `json:"month"`
		Day     int    `json:"day"`
		Shichen string `json:"shichen"`
		Gender  string `json:"gender"`
		Master  string `json:"master"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	result := logic.BaziAnalyze(req.Year, req.Month, req.Day, req.Shichen, req.Gender, req.Master)
	jsonResponse(c, 0, result)
}

func (h *FortuneHandler) DreamInterpret(c *gin.Context) {
	var req struct {
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Description == "" {
		jsonError(c, "请描述您的梦境")
		return
	}

	result := logic.InterpretDream(req.Description)
	jsonResponse(c, 0, result)
}

func (h *FortuneHandler) DivinationHexagram(c *gin.Context) {
	var req struct {
		Question string `json:"question"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "请输入所问之事")
		return
	}

	result := logic.InterpretDivination(req.Question)
	jsonResponse(c, 0, result)
}
