package services

import (
	"Hi_Tech/internal/controller"
	"Hi_Tech/internal/model"
)

func GetAllProducts() ([]model.Products, error) {
	return controller.Product.GetALlProducts(nil), nil
}
