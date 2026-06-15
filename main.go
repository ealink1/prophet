package main

import (
	"log"
	"net/http"
	"os"

	"prophet/internal/handlers"
	"prophet/internal/middleware"
	"prophet/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB("data/prophet.db")

	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// 静态文件
	r.Static("/static", "./web/static")
	r.StaticFile("/favicon.ico", "./web/static/images/favicon.ico")
	r.StaticFile("/favicon.svg", "./web/static/images/favicon.svg")
	r.StaticFile("/manifest.json", "./web/static/manifest.json")

	// API 路由
	api := r.Group("/api/v1")
	{
		// 认证
		api.POST("/auth/anonymous/init", handlers.AuthAnonymousInit)
		api.POST("/auth/me", middleware.AuthMiddleware(), handlers.AuthMe)

		// 黄历
		api.POST("/almanac/today", handlers.AlmanacToday)
		api.POST("/almanac/week", handlers.AlmanacWeek)

		// 灵签
		api.POST("/lottery/draw", handlers.LotteryDraw)

		// 八字
		api.POST("/bazi/analyze", handlers.BaziAnalyze)

		// 解梦
		api.POST("/dream/interpret", handlers.DreamInterpret)

		// 六爻
		api.POST("/divination/hexagram", handlers.DivinationHexagram)

		// 祈福灯
		api.POST("/blessing/create", middleware.AuthMiddleware(), handlers.BlessingCreate)
		api.POST("/blessing/wall", handlers.BlessingWall)
		api.POST("/blessing/stats", handlers.BlessingStats)

		// 禅修
		api.POST("/meditation/catalog", handlers.MeditationCatalog)

		// 商品
		api.POST("/products/list", handlers.ProductList)

		// 订单
		api.POST("/orders/create", middleware.AuthMiddleware(), handlers.OrderCreate)
		api.POST("/orders/list", handlers.OrderList)
		api.POST("/orders/confirm/:id", handlers.OrderConfirm)

		// 活动追踪
		api.POST("/activity/track", handlers.ActivityTrack)

		// 后台
		api.POST("/admin/dashboard", handlers.AdminDashboard)
	}

	// 页面路由 - 所有页面都返回同一个 SPA HTML
	r.NoRoute(func(c *gin.Context) {
		c.File("./web/templates/index.html")
	})

	r.GET("/", func(c *gin.Context) {
		c.File("./web/templates/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf(" server starting on :%s", port)
	r.Run(":" + port)
}
