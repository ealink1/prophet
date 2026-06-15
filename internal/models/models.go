package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	DeviceID          string `gorm:"uniqueIndex;size:64" json:"device_id"`
	Phone             string `gorm:"size:20" json:"phone"`
	Nickname          string `gorm:"size:50;default:'有缘人'" json:"nickname"`
	LuckyCode         string `gorm:"uniqueIndex;size:20" json:"lucky_code"`
	InviteCode        string `gorm:"size:20" json:"invite_code"`
	ReferredBy        uint   `json:"referred_by"`
	MeritBalance      int64  `json:" merit_balance"`
	FreeLotteryDaily  int    `gorm:"default:3" json:"free_lottery_daily"`
	FreeDreamDaily    int    `gorm:"default:5" json:"free_dream_daily"`
	FreeBaziDaily     int    `gorm:"default:1" json:"free_bazi_daily"`
	LastActiveAt      int64  `json:"last_active_at"`
}

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

type Order struct {
	gorm.Model
	OrderNo        string  `gorm:"uniqueIndex;size:40" json:"order_no"`
	UserID         uint    `json:"user_id"`
	User           User    `gorm:"foreignKey:UserID" json:"-"`
	ProductID      string  `gorm:"size:50" json:"product_id"`
	ProductName    string  `gorm:"size:100" json:"product_name"`
	Amount         float64 `json:"amount"`
	OriginalPrice  float64 `json:"original_price"`
	Status         string  `gorm:"size:20;default:'pending'" json:"status"` // pending/reviewing/paid/expired/rejected
	PaymentChannel string  `gorm:"size:20;default:'personal_qr'" json:"payment_channel"`
	ProofImage     string  `json:"proof_image"`
	ReviewNote     string  `json:"review_note"`
	Note           string  `json:"note"`
	ExpiredAt      int64   `json:"expired_at"`
}

type BlessingLamp struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	User         User   `gorm:"foreignKey:UserID" json:"-"`
	LampType     string `gorm:"size:20" json:"lamp_type"` // peace/wisdom/love/wealth
	ForPerson    string `gorm:"size:50" json:"for_person"`
	Relation     string `gorm:"size:20" json:"relation"` // father/mother/spouse/child/grandchild/friend/self
	Wish         string `gorm:"size:200" json:"wish"`
	DisplayName  string `gorm:"size:50" json:"display_name"`
	DurationHrs  int    `gorm:"default:24" json:"duration_hours"`
	Status       string `gorm:"size:20;default:'active'" json:"status"`
	ExpiresAt    int64  `json:"expires_at"`
}

type LotteryRecord struct {
	gorm.Model
	UserID         uint   `json:"user_id"`
	User           User   `gorm:"foreignKey:UserID" json:"-"`
	Master         string `gorm:"size:20" json:"master"` // huiming/mingxin/xuanzhen
	Question       string `json:"question"`
	SignNumber     string `gorm:"size:20" json:"sign_number"`
	SignTitle      string `gorm:"size:50" json:"sign_title"`
	SignLevel      string `gorm:"size:10" json:"sign_level"` // 上上/上吉/中吉/中平/下下
	SignPoem       string `json:"sign_poem"`
	SignAnalysis   string `json:"sign_analysis"`
	MasterReading  string `json:"master_reading"`
	Advice         string `json:"advice"`
	IsFree         bool   `json:"is_free"`
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
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID" json:"-"`
	BabyName    string `gorm:"size:20" json:"baby_name"`
	Gender      string `gorm:"size:10" json:"gender"`
	BirthInfo   string `json:"birth_info"`
	Style       string `gorm:"size:10" json:"style"`
	Candidates  string `json:"candidates"`
	IsCompany   bool   `json:"is_company"`
	IsFree      bool   `json:"is_free"`
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

type MeditationTrack struct {
	gorm.Model
	Title      string `gorm:"size:50" json:"title"`
	Subtitle   string `gorm:"size:100" json:"subtitle"`
	Icon       string `gorm:"size:10" json:"icon"`
	URL        string `gorm:"size:200" json:"url"`
	Duration   int    `json:"duration"`
	Genre      string `gorm:"size:20" json:"genre"`
	License    string `gorm:"size:20" json:"license"`
	Color      string `gorm:"size:10" json:"color"`
	Descript   string `json:"description"`
	SortOrder  int    `json:"sort_order"`
	IsActive   bool   `gorm:"default:true" json:"is_active"`
	PlayCount  int    `json:"play_count"`
}

type MeditationPlay struct {
	gorm.Model
	UserID          uint   `json:"user_id"`
	User            User   `gorm:"foreignKey:UserID" json:"-"`
	TrackID         uint   `json:"track_id"`
	Track           MeditationTrack `gorm:"foreignKey:TrackID" json:"-"`
	DurationListened int   `json:"duration_listened"`
	Completed       bool   `json:"completed"`
	MeritEarned     int    `json:"merit_earned"`
}

type AlmanacCache struct {
	Date      string `gorm:"primaryKey;size:10" json:"date"`
	Data      string `json:"data"`
	ExpiresAt int64  `json:"expires_at"`
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
	ReferrerID   uint   `json:"referrer_id"`
	Referrer     User   `gorm:"foreignKey:ReferrerID" json:"-"`
	ReferredID   uint   `json:"referred_id"`
	InviteCode   string `gorm:"size:20" json:"invite_code"`
	MeritReward  int64  `json:"merit_reward"`
	Status       string `gorm:"size:20;default:'pending'" json:"status"`
}

type MeritTransaction struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID" json:"-"`
	Type      string `gorm:"size:10" json:"type"` // earn/spend/withdraw/reward
	Amount    int64  `json:"amount"`
	Source    string `gorm:"size:20" json:"source"`
	RefID     uint   `json:"reference_id"`
	Note      string `json:"note"`
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
