package main

import (
	"log"
	"net/http"
	"os"

	"prophet/internal/handler"
	"prophet/internal/middleware"
	"prophet/internal/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDB("data/prophet.db")

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

	// 初始化 handlers
	authH := handler.NewAuthHandler()
	almanacH := handler.NewAlmanacHandler()
	fortuneH := handler.NewFortuneHandler()
	blessingH := handler.NewBlessingHandler()
	meditationH := handler.NewMeditationHandler()
	orderH := handler.NewOrderHandler()
	productH := handler.NewProductHandler()
	activityH := handler.NewActivityHandler()
	adminH := handler.NewAdminHandler()

	// API 路由
	api := r.Group("/api/v1")
	{
		// 认证
		api.POST("/auth/anonymous/init", authH.AnonymousInit)
		api.POST("/auth/me", middleware.AuthMiddleware(), authH.Me)

		// 黄历
		api.POST("/almanac/today", almanacH.Today)
		api.POST("/almanac/week", almanacH.Week)

		// 灵签
		api.POST("/lottery/draw", fortuneH.LotteryDraw)

		// 八字
		api.POST("/bazi/analyze", fortuneH.BaziAnalyze)

		// 解梦
		api.POST("/dream/interpret", fortuneH.DreamInterpret)

		// 六爻
		api.POST("/divination/hexagram", fortuneH.DivinationHexagram)

		// 祈福灯
		api.POST("/blessing/create", middleware.AuthMiddleware(), blessingH.Create)
		api.POST("/blessing/wall", blessingH.Wall)
		api.POST("/blessing/stats", blessingH.Stats)

		// 禅修
		api.POST("/meditation/catalog", meditationH.Catalog)

		// 商品
		api.POST("/products/list", productH.List)

		// 订单
		api.POST("/orders/create", middleware.AuthMiddleware(), orderH.Create)
		api.POST("/orders/list", orderH.List)
		api.POST("/orders/confirm/:id", orderH.Confirm)

		// 活动追踪
		api.POST("/activity/track", activityH.Track)

		// 后台
		api.POST("/admin/dashboard", adminH.Dashboard)
	}

	// 页面路由
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
