package dao

import "prophet/internal/model"

type OrderDAO struct{}

func (d *OrderDAO) Create(order *model.Order) error {
	return model.DB.Create(order).Error
}

func (d *OrderDAO) FindByID(id interface{}) (*model.Order, error) {
	var order model.Order
	if err := model.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (d *OrderDAO) FindByUserID(userID uint, status string, offset, limit int) ([]model.Order, error) {
	var orders []model.Order
	query := model.DB.Order("created_at desc")
	if status != "" {
		query = query.Where("user_id = ? AND status = ?", userID, status)
	} else {
		query = query.Where("user_id = ?", userID)
	}
	err := query.Offset(offset).Limit(limit).Find(&orders).Error
	return orders, err
}

func (d *OrderDAO) CountAll() (int64, error) {
	var total int64
	err := model.DB.Model(&model.Order{}).Count(&total).Error
	return total, err
}

func (d *OrderDAO) CountByStatus(status string) (int64, error) {
	var count int64
	err := model.DB.Model(&model.Order{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

func (d *OrderDAO) Save(order *model.Order) error {
	return model.DB.Save(order).Error
}

func (d *OrderDAO) SumPaidAmount() (float64, error) {
	var total float64
	err := model.DB.Model(&model.Order{}).Where("status = ?", "paid").Select("coalesce(sum(amount),0)").Scan(&total).Error
	return total, err
}

func (d *OrderDAO) SumTodayPaidAmount() (float64, error) {
	var total float64
	err := model.DB.Model(&model.Order{}).Where("status = ? AND date(created_at) = date('now')", "paid").Select("coalesce(sum(amount),0)").Scan(&total).Error
	return total, err
}
