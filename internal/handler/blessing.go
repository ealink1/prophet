package handler

import (
	"prophet/internal/dao"
	"prophet/internal/logic"

	"github.com/gin-gonic/gin"
)

type BlessingHandler struct {
	blessingLogic *logic.BlessingLogic
}

func NewBlessingHandler() *BlessingHandler {
	return &BlessingHandler{blessingLogic: logic.NewBlessingLogic()}
}

func (h *BlessingHandler) Create(c *gin.Context) {
	var req struct {
		LampType    string `json:"lamp_type"`
		ForPerson   string `json:"for_person"`
		Relation    string `json:"relation"`
		Wish        string `json:"wish"`
		DisplayName string `json:"display_name"`
		DurationHrs int    `json:"duration_hours"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	userID, _ := c.Get("user_id")
	lamp, err := h.blessingLogic.Create(userID.(uint), dao.CreateLampReq{
		LampType:    req.LampType,
		ForPerson:   req.ForPerson,
		Relation:    req.Relation,
		Wish:        req.Wish,
		DisplayName: req.DisplayName,
		DurationHrs: req.DurationHrs,
	})
	if err != nil {
		jsonError(c, "创建失败")
		return
	}

	jsonResponse(c, 0, gin.H{"id": lamp.ID, "message": "祈福灯已点亮"})
}

func (h *BlessingHandler) Wall(c *gin.Context) {
	views, err := h.blessingLogic.Wall()
	if err != nil {
		jsonError(c, "获取失败")
		return
	}
	jsonResponse(c, 0, gin.H{"lamps": views})
}

func (h *BlessingHandler) Stats(c *gin.Context) {
	total, todayNew, err := h.blessingLogic.Stats()
	if err != nil {
		jsonError(c, "获取失败")
		return
	}
	jsonResponse(c, 0, gin.H{"total": total, "today_new": todayNew})
}
