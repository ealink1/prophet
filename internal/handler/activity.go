package handler

import (
	"net/http"

	"prophet/internal/dao"
	"prophet/internal/model"

	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	activityDAO *dao.ActivityDAO
}

func NewActivityHandler() *ActivityHandler {
	return &ActivityHandler{activityDAO: &dao.ActivityDAO{}}
}

func (h *ActivityHandler) Track(c *gin.Context) {
	var req struct {
		DeviceID  string `json:"device_id"`
		EventType string `json:"event_type"`
		Path      string `json:"path"`
		Title     string `json:"title"`
		Referrer  string `json:"referrer"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	log := &model.ActivityLog{
		DeviceID:  req.DeviceID,
		EventType: req.EventType,
		Path:      req.Path,
		Title:     req.Title,
		Referrer:  req.Referrer,
	}
	h.activityDAO.Create(log)
	c.JSON(http.StatusOK, gin.H{"code": 0})
}
