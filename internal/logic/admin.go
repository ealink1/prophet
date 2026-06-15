package logic

import "prophet/internal/model"

type AdminLogic struct {
	userDAO  *UserDAO
	orderDAO *OrderDAO
}

type UserDAO struct{}
type OrderDAO struct{}

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

func NewAdminLogic() *AdminLogic {
	return &AdminLogic{
		userDAO:  &UserDAO{},
		orderDAO: &OrderDAO{},
	}
}

type DashboardData struct {
	TodayIncome     float64 `json:"today_income"`
	TotalIncome     float64 `json:"total_income"`
	TotalUsers      int64   `json:"total_users"`
	TodayActive     int64   `json:"today_active"`
	TotalOrders     int64   `json:"total_orders"`
	PaidOrders      int64   `json:"paid_orders"`
	PendingOrders   int64   `json:"pending_orders"`
	ReviewingOrders int64   `json:"reviewing_orders"`
}

func (l *AdminLogic) Dashboard() (*DashboardData, error) {
	data := &DashboardData{}

	data.TotalIncome, _ = l.orderDAO.SumPaidAmount()
	data.TodayIncome, _ = l.orderDAO.SumTodayPaidAmount()
	data.TotalUsers, _ = l.userDAO.CountAll()
	data.TodayActive, _ = l.userDAO.CountTodayActive()
	data.TotalOrders, _ = l.orderDAO.CountAll()
	data.PaidOrders, _ = l.orderDAO.CountByStatus("paid")
	data.PendingOrders, _ = l.orderDAO.CountByStatus("pending")
	data.ReviewingOrders, _ = l.orderDAO.CountByStatus("reviewing")

	return data, nil
}
