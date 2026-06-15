package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DeviceID         string `gorm:"uniqueIndex;size:64" json:"device_id"`
	Phone            string `gorm:"size:20" json:"phone"`
	Nickname         string `gorm:"size:50;default:'有缘人'" json:"nickname"`
	LuckyCode        string `gorm:"uniqueIndex;size:20" json:"lucky_code"`
	InviteCode       string `gorm:"size:20" json:"invite_code"`
	ReferredBy       uint   `json:"referred_by"`
	MeritBalance     int64  `json:"merit_balance"`
	FreeLotteryDaily int    `gorm:"default:3" json:"free_lottery_daily"`
	FreeDreamDaily   int    `gorm:"default:5" json:"free_dream_daily"`
	FreeBaziDaily    int    `gorm:"default:1" json:"free_bazi_daily"`
	LastActiveAt     int64  `json:"last_active_at"`
}
