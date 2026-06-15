package model

import "gorm.io/gorm"

type LotteryRecord struct {
	gorm.Model
	UserID        uint   `json:"user_id"`
	User          User   `gorm:"foreignKey:UserID" json:"-"`
	Master        string `gorm:"size:20" json:"master"`
	Question      string `json:"question"`
	SignNumber    string `gorm:"size:20" json:"sign_number"`
	SignTitle     string `gorm:"size:50" json:"sign_title"`
	SignLevel     string `gorm:"size:10" json:"sign_level"`
	SignPoem      string `json:"sign_poem"`
	SignAnalysis  string `json:"sign_analysis"`
	MasterReading string `json:"master_reading"`
	Advice        string `json:"advice"`
	IsFree        bool   `json:"is_free"`
}

type BaziRecord struct {
	gorm.Model
	UserID        uint   `json:"user_id"`
	User          User   `gorm:"foreignKey:UserID" json:"-"`
	Master        string `gorm:"size:20" json:"master"`
	BirthYear     int    `json:"birth_year"`
	BirthMonth    int    `json:"birth_month"`
	BirthDay      int    `json:"birth_day"`
	BirthShichen  string `gorm:"size:10" json:"birth_shichen"`
	Gender        string `gorm:"size:10" json:"gender"`
	BaziResult    string `json:"bazi_result"`
	Personality   string `json:"personality"`
	Career        string `json:"career"`
	Wealth        string `json:"wealth"`
	Relationship  string `json:"relationship"`
	Health        string `json:"health"`
	MasterReading string `json:"master_reading"`
	IsFree        bool   `json:"is_free"`
}

type DreamRecord struct {
	gorm.Model
	UserID           uint   `json:"user_id"`
	User             User   `gorm:"foreignKey:UserID" json:"-"`
	DreamDescription string `json:"dream_description"`
	Interpretation   string `json:"interpretation"`
	LuckyLevel       string `gorm:"size:10" json:"lucky_level"`
	Advice           string `json:"advice"`
	IsFree           bool   `json:"is_free"`
}

type DivinationRecord struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	User         User   `gorm:"foreignKey:UserID" json:"-"`
	Question     string `json:"question"`
	Hexagram     string `json:"hexagram"`
	HexagramName string `gorm:"size:20" json:"hexagram_name"`
	Analysis     string `json:"analysis"`
	IsFree       bool   `json:"is_free"`
}
