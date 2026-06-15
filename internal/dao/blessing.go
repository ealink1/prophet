package dao

import (
	"time"

	"prophet/internal/model"
)

type BlessingDAO struct{}

func (d *BlessingDAO) Create(lamp *model.BlessingLamp) error {
	return model.DB.Create(lamp).Error
}

func (d *BlessingDAO) FindActiveWall(limit int) ([]model.BlessingLamp, error) {
	var lamps []model.BlessingLamp
	err := model.DB.Where("status = ?", "active").Order("created_at desc").Limit(limit).Find(&lamps).Error
	return lamps, err
}

func (d *BlessingDAO) CountAll() (int64, error) {
	var total int64
	err := model.DB.Model(&model.BlessingLamp{}).Count(&total).Error
	return total, err
}

func (d *BlessingDAO) CountToday() (int64, error) {
	var count int64
	err := model.DB.Model(&model.BlessingLamp{}).Where("date(created_at) = date('now')").Count(&count).Error
	return count, err
}

func NewBlessingLamp(userID uint, req CreateLampReq) *model.BlessingLamp {
	return &model.BlessingLamp{
		UserID:      userID,
		LampType:    req.LampType,
		ForPerson:   req.ForPerson,
		Relation:    req.Relation,
		Wish:        req.Wish,
		DisplayName: req.DisplayName,
		DurationHrs: req.DurationHrs,
		Status:      "active",
		ExpiresAt:   time.Now().Add(time.Duration(req.DurationHrs) * time.Hour).Unix(),
	}
}

type CreateLampReq struct {
	LampType    string
	ForPerson   string
	Relation    string
	Wish        string
	DisplayName string
	DurationHrs int
}
