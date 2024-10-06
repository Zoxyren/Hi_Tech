package controller

import (
	"Hi_Tech/internal/model"
	"Hi_Tech/internal/repository"
	"Hi_Tech/internal/services"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Product struct {
	products *[]model.Products
	db       *sql.DB
}

func (p *Product) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	productRepo := repository.ProductRepository{Db: p.db}
	productService := services.ProductService{Repo: productRepo, Db: p.db}
	products, err := productService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
