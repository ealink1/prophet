package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID     string  `gorm:"uniqueIndex;size:50" json:"product_id"`
	Name          string  `gorm:"size:100" json:"name"`
	Category      string  `gorm:"size:30" json:"category"`
	Description   string  `json:"description"`
	OriginalPrice float64 `json:"original_price"`
	Price         float64 `json:"price"`
	Badge         string  `gorm:"size:20" json:"badge"`
	Benefits      string  `json:"benefits"`
	IsActive      bool    `gorm:"default:true" json:"is_active"`
	SortOrder     int     `json:"sort_order"`
}
