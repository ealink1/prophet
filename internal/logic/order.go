package logic

import (
	"time"

	"prophet/internal/dao"
	"prophet/internal/model"
)

type OrderLogic struct {
	orderDAO   *dao.OrderDAO
	productDAO *dao.ProductDAO
}

func NewOrderLogic() *OrderLogic {
	return &OrderLogic{
		orderDAO:   &dao.OrderDAO{},
		productDAO: &dao.ProductDAO{},
	}
}

func (l *OrderLogic) Create(userID uint, productID string) (*model.Order, error) {
	product, err := l.productDAO.FindByProductID(productID)
	if err != nil {
		return nil, err
	}

	order := &model.Order{
		OrderNo:       GenerateOrderNo(),
		UserID:        userID,
		ProductID:     product.ProductID,
		ProductName:   product.Name,
		Amount:        product.Price,
		OriginalPrice: product.OriginalPrice,
		Status:        "pending",
		ExpiredAt:     time.Now().Add(30 * time.Minute).Unix(),
	}
	if err := l.orderDAO.Create(order); err != nil {
		return nil, err
	}
	return order, nil
}

func (l *OrderLogic) List(userID uint, status string, page, limit int) ([]model.Order, int64, error) {
	offset := (page - 1) * limit
	orders, err := l.orderDAO.FindByUserID(userID, status, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	total, err := l.orderDAO.CountAll()
	if err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}

func (l *OrderLogic) Confirm(id interface{}) error {
	order, err := l.orderDAO.FindByID(id)
	if err != nil {
		return err
	}
	order.Status = "paid"
	return l.orderDAO.Save(order)
}
