package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 40101, "message": "未登录"})
			c.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		// 简单的 token 格式: token_{user_id}_{random}
		// 从 token 中提取 user_id
		parts := strings.Split(token, "_")
		if len(parts) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 40101, "message": "无效token"})
			c.Abort()
			return
		}

		// 尝试解析 user_id
		var userID uint
		for _, ch := range parts[1] {
			if ch >= '0' && ch <= '9' {
				userID = userID*10 + uint(ch-'0')
			}
		}
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 40101, "message": "无效token"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}

func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "" {
			token = strings.TrimPrefix(token, "Bearer ")
			parts := strings.Split(token, "_")
			if len(parts) >= 2 {
				var userID uint
				for _, ch := range parts[1] {
					if ch >= '0' && ch <= '9' {
						userID = userID*10 + uint(ch-'0')
					}
				}
				if userID > 0 {
					c.Set("user_id", userID)
				}
			}
		}
		c.Next()
	}
}
