package model

import "gorm.io/gorm"

type MeditationTrack struct {
	gorm.Model
	Title     string `gorm:"size:50" json:"title"`
	Subtitle  string `gorm:"size:100" json:"subtitle"`
	Icon      string `gorm:"size:10" json:"icon"`
	URL       string `gorm:"size:200" json:"url"`
	Duration  int    `json:"duration"`
	Genre     string `gorm:"size:20" json:"genre"`
	License   string `gorm:"size:20" json:"license"`
	Color     string `gorm:"size:10" json:"color"`
	Descript  string `json:"description"`
	SortOrder int    `json:"sort_order"`
	IsActive  bool   `gorm:"default:true" json:"is_active"`
	PlayCount int    `json:"play_count"`
}

type MeditationPlay struct {
	gorm.Model
	UserID           uint            `json:"user_id"`
	User             User            `gorm:"foreignKey:UserID" json:"-"`
	TrackID          uint            `json:"track_id"`
	Track            MeditationTrack `gorm:"foreignKey:TrackID" json:"-"`
	DurationListened int             `json:"duration_listened"`
	Completed        bool            `json:"completed"`
	MeritEarned      int             `json:"merit_earned"`
}
