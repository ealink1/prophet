package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbPath string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB.AutoMigrate(
		&User{}, &Product{}, &Order{}, &BlessingLamp{},
		&LotteryRecord{}, &BaziRecord{}, &DreamRecord{},
		&PalmistryRecord{}, &NamingRecord{}, &DivinationRecord{},
		&MeditationTrack{}, &MeditationPlay{}, &AlmanacCache{},
		&ActivityLog{}, &UserHistory{}, &Referral{},
		&MeritTransaction{}, &AdminUser{}, &AuditLog{}, &SystemConfig{},
	)

	seedProducts()
	seedMeditationTracks()
	seedSystemConfig()
}

func seedProducts() {
	var count int64
	DB.Model(&Product{}).Count(&count)
	if count > 0 {
		return
	}
	products := []Product{
		{ProductID: "blessing_lamp", Name: "祈福供灯", Category: "blessing", OriginalPrice: 6.60, Price: 6.60, Benefits: `["点亮功德灯","供奉指定时长","功德灯墙展示"]`},
		{ProductID: "blessing_peace", Name: "祈福供灯·平安灯", Category: "blessing", OriginalPrice: 3.90, Price: 3.72, Benefits: `["平安灯","供奉指定时长"]`},
		{ProductID: "blessing_wisdom", Name: "祈福供灯·智慧灯", Category: "blessing", OriginalPrice: 9.90, Price: 9.80, Benefits: `["智慧灯","供奉指定时长"]`},
		{ProductID: "blessing_love", Name: "祈福供灯·姻缘灯", Category: "blessing", OriginalPrice: 5.90, Price: 5.71, Benefits: `["姻缘灯","供奉指定时长"]`},
		{ProductID: "blessing_wealth", Name: "祈福供灯·财福灯", Category: "blessing", OriginalPrice: 3.90, Price: 3.82, Benefits: `["财福灯","供奉指定时长"]`},
		{ProductID: "extra_lottery", Name: "关帝灵签·加抽一次", Category: "lottery", OriginalPrice: 2.90, Price: 2.86, Benefits: `["加抽一次灵签","师父开示"]`},
		{ProductID: "extra_dream", Name: "周公解梦·加抽一次", Category: "dream", OriginalPrice: 3.90, Price: 3.78, Benefits: `["加抽一次解梦"]`},
		{ProductID: "unlock_bazi", Name: "八字精批·完整解读", Category: "bazi", OriginalPrice: 3.90, Price: 3.71, Benefits: `["完整十神/大运/流年","古籍引用"]`},
		{ProductID: "single_liunian", Name: "流年运势详批", Category: "bazi", OriginalPrice: 9.90, Price: 9.90, Badge: "首单特惠", Benefits: `["12月逐月运势","贵人/桃花/财禄提示","师父开示"]`},
		{ProductID: "single_bazi_deep", Name: "八字精批深度版", Category: "bazi", OriginalPrice: 19.90, Price: 19.90, Badge: "明星产品", Benefits: `["完整十神/大运/流年","古籍引用","PDF报告"]`},
		{ProductID: "unlock_palmistry", Name: "手相图解·完整解读", Category: "palmistry", OriginalPrice: 6.60, Price: 6.45, Benefits: `["拍照上传","手纹细看","手纹命理详解"]`},
		{ProductID: "unlock_naming", Name: "宝宝起名·解锁全部30名", Category: "naming", OriginalPrice: 66.00, Price: 65.97, Benefits: `["30个候选名","音韵/笔画/五行评分","典故出处"]`},
		{ProductID: "single_naming_pro", Name: "宝宝起名VIP", Category: "naming", OriginalPrice: 49.90, Price: 49.90, Badge: "热销", Benefits: `["30个候选名","音韵/笔画/五行评分","典故出处"]`},
		{ProductID: "single_company", Name: "公司起名", Category: "naming", OriginalPrice: 99.90, Price: 99.90, Badge: "企业版", Benefits: `["行业五行匹配","品牌寓意","5个候选方案"]`},
		{ProductID: "extra_divination", Name: "六爻占卜·加抽一次", Category: "divination", OriginalPrice: 2.90, Price: 2.84, Benefits: `["加抽一次六爻"]`},
		{ProductID: "single_hehun", Name: "两人合婚", Category: "bazi", OriginalPrice: 29.90, Price: 29.90, Badge: "情感推荐", Benefits: `["双方八字配对","五行互补分析","古籍参考"]`},
	}
	DB.Create(&products)
}

func seedMeditationTracks() {
	var count int64
	DB.Model(&MeditationTrack{}).Count(&count)
	if count > 0 {
		return
	}
	tracks := []MeditationTrack{
		{Title: "菩提苑主题曲", Subtitle: "金光普照·寺院庄严", Icon: "🧘", URL: "/static/audio/bodhi_theme.mp3", Duration: 177, Genre: "主题", License: "项目原创", Color: "#c9a05c", Descript: "菩提苑专属主题音乐，金光普照感，作为开场冥想最合适。", SortOrder: 1},
		{Title: "菩提苑", Subtitle: "苑中清雅·万缘澄定", Icon: "🌿", URL: "/static/audio/bodhi_garden.mp3", Duration: 171, Genre: "禅意", License: "项目原创", Color: "#7BA686", Descript: "走在菩提苑的小径上，万缘澄定，最适合午后冥想。", SortOrder: 2},
		{Title: "菩提苑·轻音乐", Subtitle: "轻柔禅意·心境清明", Icon: "🧘", URL: "/static/audio/bodhi_light.mp3", Duration: 195, Genre: "轻禅", License: "项目原创", Color: "#D4BC8A", Descript: "轻柔版菩提苑主题，更适合长时间禅修陪伴。", SortOrder: 3},
		{Title: "菩提苑·渡尘缘", Subtitle: "渡过尘缘·返照本心", Icon: "🛤️", URL: "/static/audio/bodhi_crossing.mp3", Duration: 219, Genre: "禅悟", License: "项目原创", Color: "#3D5A80", Descript: "象征渡过红尘缠缚，回归本心清净的旋律。", SortOrder: 4},
		{Title: "宝殿晨曦", Subtitle: "晨钟初响·佛光初临", Icon: "🌅", URL: "/static/audio/palace_dawn.mp3", Duration: 168, Genre: "晨修", License: "项目原创", Color: "#E89B5C", Descript: "晨曦时分宝殿初开的庄严景象，最适合清晨第一坐。", SortOrder: 5},
		{Title: "禅坐", Subtitle: "结跏趺坐·身心安住", Icon: "🧘", URL: "/static/audio/zen_sit.mp3", Duration: 156, Genre: "正念", License: "项目原创", Color: "#5A7C65", Descript: "纯净的禅坐音乐，引导身心快速安住于当下。", SortOrder: 6},
		{Title: "禅意", Subtitle: "万像皆禅·处处是道场", Icon: "☯️", URL: "/static/audio/zen_mind.mp3", Duration: 192, Genre: "禅意", License: "项目原创", Color: "#8A8A8A", Descript: "在喧嚣中也能听见的内心静默，让禅意流淌。", SortOrder: 7},
		{Title: "琉璃月", Subtitle: "月光琉璃·照见五蕴", Icon: "🌙", URL: "/static/audio/crystal_moon.mp3", Duration: 211, Genre: "禅韵", License: "项目原创", Color: "#a3c5ab", Descript: "如月光透过琉璃般清澄的旋律，照见五蕴皆空。", SortOrder: 8},
		{Title: "大悲咒", Subtitle: "观音大悲·消业除障", Icon: "🙏", URL: "/static/audio/great_compassion.mp3", Duration: 246, Genre: "梵音", License: "传统佛曲", Color: "#c43d3d", Descript: "观世音菩萨大悲咒梵音版，消业除障、增长慈悲。", SortOrder: 9},
		{Title: "心经", Subtitle: "般若智慧·照见空性", Icon: "📿", URL: "/static/audio/heart_sutra.mp3", Duration: 235, Genre: "梵音", License: "传统佛曲", Color: "#c9a05c", Descript: "《般若波罗蜜多心经》260字浓缩般若智慧。", SortOrder: 10},
	}
	DB.Create(&tracks)
}

func seedSystemConfig() {
	var count int64
	DB.Model(&SystemConfig{}).Count(&count)
	if count > 0 {
		return
	}
	configs := []SystemConfig{
		{Key: "site_name", Value: "菩提苑", Description: "站点名称"},
		{Key: "free_lottery_daily", Value: "3", Description: "每日免费灵签次数"},
		{Key: "free_dream_daily", Value: "5", Description: "每日免费解梦次数"},
		{Key: "free_bazi_daily", Value: "1", Description: "每日免费八字次数"},
		{Key: "blessing_default_price", Value: "6.60", Description: "祈福灯默认价格"},
		{Key: "order_expire_minutes", Value: "30", Description: "订单超时时间"},
	}
	DB.Create(&configs)
}
