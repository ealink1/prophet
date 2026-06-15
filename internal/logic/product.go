package logic

import (
	"prophet/internal/dao"
)

type ProductLogic struct {
	productDAO *dao.ProductDAO
}

func NewProductLogic() *ProductLogic {
	return &ProductLogic{productDAO: &dao.ProductDAO{}}
}

func ListActiveProducts() (interface{}, error) {
	return (&ProductLogic{productDAO: &dao.ProductDAO{}}).productDAO.FindActive()
}
