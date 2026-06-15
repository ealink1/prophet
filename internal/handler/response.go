package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func jsonResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": code, "data": data})
}

func jsonError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": -1, "message": msg})
}
