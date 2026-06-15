package model

import "gorm.io/gorm"

type PalmistryRecord struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	User         User   `gorm:"foreignKey:UserID" json:"-"`
	ImageURL     string `json:"image_url"`
	PalmAnalysis string `json:"palm_analysis"`
	Reading      string `json:"reading"`
	IsFree       bool   `json:"is_free"`
}

type NamingRecord struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	User       User   `gorm:"foreignKey:UserID" json:"-"`
	BabyName   string `gorm:"size:20" json:"baby_name"`
	Gender     string `gorm:"size:10" json:"gender"`
	BirthInfo  string `json:"birth_info"`
	Style      string `gorm:"size:10" json:"style"`
	Candidates string `json:"candidates"`
	IsCompany  bool   `json:"is_company"`
	IsFree     bool   `json:"is_free"`
}

type UserHistory struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"-"`
	Kind     string `gorm:"size:20" json:"kind"`
	Title    string `gorm:"size:50" json:"title"`
	Subtitle string `gorm:"size:100" json:"subtitle"`
	Payload  string `json:"payload"`
}

type Referral struct {
	gorm.Model
	ReferrerID  uint   `json:"referrer_id"`
	Referrer    User   `gorm:"foreignKey:ReferrerID" json:"-"`
	ReferredID  uint   `json:"referred_id"`
	InviteCode  string `gorm:"size:20" json:"invite_code"`
	MeritReward int64  `json:"merit_reward"`
	Status      string `gorm:"size:20;default:'pending'" json:"status"`
}

type MeritTransaction struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
	Type   string `gorm:"size:10" json:"type"`
	Amount int64  `json:"amount"`
	Source string `gorm:"size:20" json:"source"`
	RefID  uint   `json:"reference_id"`
	Note   string `json:"note"`
}

type AdminUser struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;size:50" json:"username"`
	PasswordHash string `gorm:"size:200" json:"-"`
	Role         string `gorm:"size:20;default:'admin'" json:"role"`
	IsActive     bool   `gorm:"default:true" json:"is_active"`
}

type AuditLog struct {
	gorm.Model
	AdminID    uint   `json:"admin_id"`
	Action     string `gorm:"size:30" json:"action"`
	TargetType string `gorm:"size:20" json:"target_type"`
	TargetID   uint   `json:"target_id"`
	Detail     string `json:"detail"`
	IPAddress  string `gorm:"size:50" json:"ip_address"`
}

type SystemConfig struct {
	Key         string `gorm:"primaryKey;size:50" json:"key"`
	Value       string `json:"value"`
	Description string `gorm:"size:200" json:"description"`
}

type ActivityLog struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	DeviceID  string `gorm:"size:64" json:"device_id"`
	EventType string `gorm:"size:20" json:"event_type"`
	Path      string `gorm:"size:100" json:"path"`
	Title     string `gorm:"size:50" json:"title"`
	Referrer  string `json:"referrer"`
}

type AlmanacCache struct {
	Date      string `gorm:"primaryKey;size:10" json:"date"`
	Data      string `json:"data"`
	ExpiresAt int64  `json:"expires_at"`
}
