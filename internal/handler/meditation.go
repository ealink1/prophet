package handler

import (
	"prophet/internal/logic"

	"github.com/gin-gonic/gin"
)

type MeditationHandler struct {
	meditationLogic *logic.MeditationLogic
}

func NewMeditationHandler() *MeditationHandler {
	return &MeditationHandler{meditationLogic: logic.NewMeditationLogic()}
}

func (h *MeditationHandler) Catalog(c *gin.Context) {
	tracks, guided, quote := h.meditationLogic.Catalog()
	jsonResponse(c, 0, gin.H{
		"tracks": tracks,
		"guided": guided,
		"quote":  quote,
	})
}
