package dao

import "prophet/internal/model"

type UserDAO struct{}

func (d *UserDAO) FindByDeviceID(deviceID string) (*model.User, error) {
	var user model.User
	result := model.DB.Where("device_id = ?", deviceID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (d *UserDAO) FindByID(id interface{}) (*model.User, error) {
	var user model.User
	if err := model.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDAO) Create(user *model.User) error {
	return model.DB.Create(user).Error
}

func (d *UserDAO) CountAll() (int64, error) {
	var total int64
	err := model.DB.Model(&model.User{}).Count(&total).Error
	return total, err
}

func (d *UserDAO) CountTodayActive() (int64, error) {
	var count int64
	err := model.DB.Model(&model.User{}).Where("date(last_active_at) = date('now')").Count(&count).Error
	return count, err
}
