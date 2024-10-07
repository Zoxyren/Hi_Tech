package services

import (
	"Hi_Tech/internal/model"
	"Hi_Tech/internal/repository"
	"database/sql"
)

type ProductService struct {
	Repo repository.ProductRepository
	Db   *sql.DB
}

func (s *ProductService) GetAllProducts() ([]model.Products, error) {
	return s.Repo.FetchAll()
}

func (s *ProductService) GetProductById(id int) (*model.Products, error) {
	return s.Repo.FetchProductById(id)
}

func (s *ProductService) AddProduct(product *model.Products) error {
	return s.Repo.InsertItem(product)
}
