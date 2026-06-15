package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"prophet/internal/models"
	"prophet/internal/services"

	"github.com/gin-gonic/gin"
)

func jsonResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": code, "data": data})
}

func jsonError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": -1, "message": msg})
}

func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func generateOrderNo() string {
	return fmt.Sprintf("ord_%s_%s", time.Now().Format("20060102150405"), generateID()[:8])
}

// ==================== Auth ====================

func AuthAnonymousInit(c *gin.Context) {
	var req struct {
		DeviceID string `json:"device_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	var user models.User
	result := models.DB.Where("device_id = ?", req.DeviceID).First(&user)
	if result.Error != nil {
		luckyCode := fmt.Sprintf("佛缘%04d", time.Now().UnixNano()%10000)
		user = models.User{
			DeviceID:  req.DeviceID,
			Nickname:  "善信" + luckyCode,
			LuckyCode: luckyCode,
		}
		models.DB.Create(&user)
	}

	token := fmt.Sprintf("token_%s_%s", user.ID, generateID()[:16])
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func AuthMe(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var user models.User
	models.DB.First(&user, userID)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"user": user}})
}

// ==================== Almanac ====================

func AlmanacToday(c *gin.Context) {
	var req struct {
		Date string `json:"date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Date == "" {
		req.Date = time.Now().Format("2006-01-02")
	}

	data := services.ComputeAlmanac(req.Date)
	jsonResponse(c, 0, data)
}

func AlmanacWeek(c *gin.Context) {
	data := services.ComputeWeekAlmanac()
	jsonResponse(c, 0, gin.H{"items": data})
}

// ==================== Lottery ====================

func LotteryDraw(c *gin.Context) {
	var req struct {
		Master   string `json:"master"`
		Question string `json:"question"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	result := services.DrawLottery(req.Master)
	jsonResponse(c, 0, result)
}

// ==================== Bazi ====================

func BaziAnalyze(c *gin.Context) {
	var req struct {
		Year     int    `json:"year"`
		Month    int    `json:"month"`
		Day      int    `json:"day"`
		Shichen  string `json:"shichen"`
		Gender   string `json:"gender"`
		Master   string `json:"master"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	result := services.BaziAnalyze(req.Year, req.Month, req.Day, req.Shichen, req.Gender, req.Master)
	jsonResponse(c, 0, result)
}

// ==================== Dream ====================

func DreamInterpret(c *gin.Context) {
	var req struct {
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Description == "" {
		jsonError(c, "请描述您的梦境")
		return
	}

	result := services.InterpretDream(req.Description)
	jsonResponse(c, 0, result)
}

// ==================== Divination ====================

func DivinationHexagram(c *gin.Context) {
	var req struct {
		Question string `json:"question"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "请输入所问之事")
		return
	}

	result := services.InterpretDivination(req.Question)
	jsonResponse(c, 0, result)
}

// ==================== Blessing Lamps ====================

func BlessingCreate(c *gin.Context) {
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
	lamp := models.BlessingLamp{
		UserID:      userID.(uint),
		LampType:    req.LampType,
		ForPerson:   req.ForPerson,
		Relation:    req.Relation,
		Wish:        req.Wish,
		DisplayName: req.DisplayName,
		DurationHrs: req.DurationHrs,
		Status:      "active",
		ExpiresAt:   time.Now().Add(time.Duration(req.DurationHrs) * time.Hour).Unix(),
	}
	models.DB.Create(&lamp)
	jsonResponse(c, 0, gin.H{"id": lamp.ID, "message": "祈福灯已点亮"})
}

func BlessingWall(c *gin.Context) {
	var lamps []models.BlessingLamp
	models.DB.Where("status = ?", "active").Order("created_at desc").Limit(50).Find(&lamps)

	type LampView struct {
		DisplayName string `json:"display_name"`
		Relation    string `json:"relation"`
		LampType    string `json:"lamp_type"`
		Wish        string `json:"wish"`
		LitAt       string `json:"lit_at"`
	}

	var views []LampView
	for _, l := range lamps {
		name := l.DisplayName
		if len([]rune(name)) > 1 {
			name = string([]rune(name)[:1]) + "**"
		}
		views = append(views, LampView{
			DisplayName: name,
			Relation:    l.Relation,
			LampType:    l.LampType,
			Wish:        l.Wish,
			LitAt:       l.CreatedAt.Format("2006-01-02 15:04"),
		})
	}
	jsonResponse(c, 0, gin.H{"lamps": views})
}

func BlessingStats(c *gin.Context) {
	var total int64
	models.DB.Model(&models.BlessingLamp{}).Count(&total)
	var today int64
	models.DB.Model(&models.BlessingLamp{}).Where("date(created_at) = date('now')").Count(&today)
	jsonResponse(c, 0, gin.H{"total": total, "today_new": today})
}

// ==================== Meditation ====================

func MeditationCatalog(c *gin.Context) {
	var tracks []models.MeditationTrack
	models.DB.Where("is_active = ?", true).Order("sort_order").Find(&tracks)

	guided := []map[string]interface{}{
		{"id": "10min", "title": "十分钟入门", "subtitle": "适合初学者", "duration": 600,
			"steps": []string{"盘腿端坐，背挺直", "深呼吸三次，吸气数4秒，呼气数6秒", "把注意力放在鼻尖呼吸的进出", "杂念升起时不评判，温柔回到呼吸", "结束时双手合掌，回向众生"}},
		{"id": "20min", "title": "二十分钟正念", "subtitle": "进阶练习", "duration": 1200,
			"steps": []string{"三下吐纳调息", "观呼吸：注意力锁定鼻尖出入气", "扫描身体：从头顶到脚趾，依次放松每一处", "观念头来去：见妄念升起即知见，不跟随", "回向：愿一切众生离苦得乐"}},
	}

	quote := map[string]string{
		"text":   "不忘初心，方得始终",
		"source": "《华严经》",
	}

	jsonResponse(c, 0, gin.H{
		"tracks":  tracks,
		"guided":  guided,
		"quote":   quote,
	})
}

// ==================== Orders ====================

func OrderCreate(c *gin.Context) {
	var req struct {
		ProductID string `json:"product_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		jsonError(c, "参数错误")
		return
	}

	var product models.Product
	if result := models.DB.Where("product_id = ?", req.ProductID).First(&product); result.Error != nil {
		jsonError(c, "商品不存在")
		return
	}

	userID, _ := c.Get("user_id")
	order := models.Order{
		OrderNo:       generateOrderNo(),
		UserID:        userID.(uint),
		ProductID:     product.ProductID,
		ProductName:   product.Name,
		Amount:        product.Price,
		OriginalPrice: product.OriginalPrice,
		Status:        "pending",
		ExpiredAt:     time.Now().Add(30 * time.Minute).Unix(),
	}
	models.DB.Create(&order)
	jsonResponse(c, 0, order)
}

func OrderList(c *gin.Context) {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	var orders []models.Order
	query := models.DB.Order("created_at desc")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	query.Offset(offset).Limit(limit).Find(&orders)

	var total int64
	models.DB.Model(&models.Order{}).Count(&total)
	jsonResponse(c, 0, gin.H{"orders": orders, "total": total})
}

func OrderConfirm(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if result := models.DB.First(&order, id); result.Error != nil {
		jsonError(c, "订单不存在")
		return
	}
	order.Status = "paid"
	models.DB.Save(&order)
	jsonResponse(c, 0, gin.H{"message": "确认成功"})
}

// ==================== Products ====================

func ProductList(c *gin.Context) {
	var products []models.Product
	models.DB.Where("is_active = ?", true).Order("sort_order").Find(&products)
	jsonResponse(c, 0, products)
}

// ==================== Activity ====================

func ActivityTrack(c *gin.Context) {
	var req struct {
		DeviceID string `json:"device_id"`
		EventType string `json:"event_type"`
		Path     string `json:"path"`
		Title    string `json:"title"`
		Referrer string `json:"referrer"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	log := models.ActivityLog{
		DeviceID:  req.DeviceID,
		EventType: req.EventType,
		Path:      req.Path,
		Title:     req.Title,
		Referrer:  req.Referrer,
	}
	models.DB.Create(&log)
	c.JSON(http.StatusOK, gin.H{"code": 0})
}

// ==================== Admin ====================

func AdminDashboard(c *gin.Context) {
	var totalIncome float64
	models.DB.Model(&models.Order{}).Where("status = ?", "paid").Select("coalesce(sum(amount),0)").Scan(&totalIncome)

	var todayIncome float64
	models.DB.Model(&models.Order{}).Where("status = ? AND date(created_at) = date('now')", "paid").Select("coalesce(sum(amount),0)").Scan(&todayIncome)

	var totalUsers int64
	models.DB.Model(&models.User{}).Count(&totalUsers)
	var todayActive int64
	models.DB.Model(&models.User{}).Where("date(last_active_at) = date('now')").Count(&todayActive)

	var totalOrders, paidOrders, pendingOrders, reviewingOrders int64
	models.DB.Model(&models.Order{}).Count(&totalOrders)
	models.DB.Model(&models.Order{}).Where("status = ?", "paid").Count(&paidOrders)
	models.DB.Model(&models.Order{}).Where("status = ?", "pending").Count(&pendingOrders)
	models.DB.Model(&models.Order{}).Where("status = ?", "reviewing").Count(&reviewingOrders)

	jsonResponse(c, 0, gin.H{
		"today_income":  todayIncome,
		"total_income":  totalIncome,
		"total_users":   totalUsers,
		"today_active":  todayActive,
		"total_orders":  totalOrders,
		"paid_orders":   paidOrders,
		"pending_orders": pendingOrders,
		"reviewing_orders": reviewingOrders,
	})
}
