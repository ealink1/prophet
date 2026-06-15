package dao

import "prophet/internal/model"

type ActivityDAO struct{}

func (d *ActivityDAO) Create(log *model.ActivityLog) error {
	return model.DB.Create(log).Error
}
