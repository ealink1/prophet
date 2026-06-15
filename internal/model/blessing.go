package model

import "gorm.io/gorm"

type BlessingLamp struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"-"`
	LampType    string `gorm:"size:20" json:"lamp_type"`
	ForPerson   string `gorm:"size:50" json:"for_person"`
	Relation    string `gorm:"size:20" json:"relation"`
	Wish        string `gorm:"size:200" json:"wish"`
	DisplayName string `gorm:"size:50" json:"display_name"`
	DurationHrs int    `gorm:"default:24" json:"duration_hours"`
	Status      string `gorm:"size:20;default:'active'" json:"status"`
	ExpiresAt   int64  `json:"expires_at"`
}
