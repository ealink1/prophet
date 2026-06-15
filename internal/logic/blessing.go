package logic

import (
	"prophet/internal/dao"
	"prophet/internal/model"
)

type BlessingLogic struct {
	blessingDAO *dao.BlessingDAO
}

func NewBlessingLogic() *BlessingLogic {
	return &BlessingLogic{blessingDAO: &dao.BlessingDAO{}}
}

type BlessingLampView struct {
	DisplayName string `json:"display_name"`
	Relation    string `json:"relation"`
	LampType    string `json:"lamp_type"`
	Wish        string `json:"wish"`
	LitAt       string `json:"lit_at"`
}

func (l *BlessingLogic) Create(userID uint, req dao.CreateLampReq) (*model.BlessingLamp, error) {
	lamp := dao.NewBlessingLamp(userID, req)
	if err := l.blessingDAO.Create(lamp); err != nil {
		return nil, err
	}
	return lamp, nil
}

func (l *BlessingLogic) Wall() ([]BlessingLampView, error) {
	lamps, err := l.blessingDAO.FindActiveWall(50)
	if err != nil {
		return nil, err
	}

	var views []BlessingLampView
	for _, lamp := range lamps {
		name := lamp.DisplayName
		if len([]rune(name)) > 1 {
			name = string([]rune(name)[:1]) + "**"
		}
		views = append(views, BlessingLampView{
			DisplayName: name,
			Relation:    lamp.Relation,
			LampType:    lamp.LampType,
			Wish:        lamp.Wish,
			LitAt:       lamp.CreatedAt.Format("2006-01-02 15:04"),
		})
	}
	return views, nil
}

func (l *BlessingLogic) Stats() (total int64, todayNew int64, err error) {
	total, err = l.blessingDAO.CountAll()
	if err != nil {
		return
	}
	todayNew, err = l.blessingDAO.CountToday()
	return
}
