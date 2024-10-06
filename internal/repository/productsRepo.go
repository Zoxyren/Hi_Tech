package repository

import (
	"Hi_Tech/internal/errorHandling"
	"Hi_Tech/internal/model"
)

var products []model.Products

func GetAllItems() []model.Products {
	return products

}
func GetProductByID(id int) (*model.Products, error) {
	for _, p := range products {
		if p.ProductID == id {
			return &p, nil
		}
	}
	return nil, errorHandling.ErrItemNotFound
}
