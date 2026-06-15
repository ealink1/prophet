package dao

import "prophet/internal/model"

type ProductDAO struct{}

func (d *ProductDAO) FindByProductID(productID string) (*model.Product, error) {
	var product model.Product
	result := model.DB.Where("product_id = ?", productID).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (d *ProductDAO) FindActive() ([]model.Product, error) {
	var products []model.Product
	err := model.DB.Where("is_active = ?", true).Order("sort_order").Find(&products).Error
	return products, err
}
