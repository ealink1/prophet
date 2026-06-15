package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderNo        string  `gorm:"uniqueIndex;size:40" json:"order_no"`
	UserID         uint    `json:"user_id"`
	User           User    `gorm:"foreignKey:UserID" json:"-"`
	ProductID      string  `gorm:"size:50" json:"product_id"`
	ProductName    string  `gorm:"size:100" json:"product_name"`
	Amount         float64 `json:"amount"`
	OriginalPrice  float64 `json:"original_price"`
	Status         string  `gorm:"size:20;default:'pending'" json:"status"`
	PaymentChannel string  `gorm:"size:20;default:'personal_qr'" json:"payment_channel"`
	ProofImage     string  `json:"proof_image"`
	ReviewNote     string  `json:"review_note"`
	Note           string  `json:"note"`
	ExpiredAt      int64   `json:"expired_at"`
}
