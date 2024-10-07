package controller

import (
	"Hi_Tech/internal/errorHandling"
	"Hi_Tech/internal/model"
	"Hi_Tech/internal/repository"
	"Hi_Tech/internal/services"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Product struct {
	products *[]model.Products
	db       *sql.DB
}
type ErrorResponse struct {
	error *errorHandling.ErrorResponse
}

func (p *Product) GetAllProducts(w http.ResponseWriter, r *http.Request) error {
	productRepo := repository.ProductRepository{Db: p.db}
	productService := services.ProductService{Repo: productRepo, Db: p.db}
	products, err := productService.GetAllProducts()
	if err != nil {
		return errorHandling.ErrInternalServer
	}
	json.NewEncoder(w).Encode(products)
	return nil
}

func (p *Product) GetProductById(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return errorHandling.ErrItemWithIDNotFound
	}

	productRepo := repository.ProductRepository{Db: p.db}
	productService := services.ProductService{Repo: productRepo, Db: p.db}
	product, err := productService.GetProductById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorHandling.ErrItemsNotFound
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return nil
	}
	json.NewEncoder(w).Encode(product)
	return nil
}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) error {
	var product model.Products
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return errorHandling.ErrorAddingProduct

	}

	productRepo := repository.ProductRepository{Db: p.db}
	productService := services.ProductService{Repo: productRepo, Db: p.db}
	err = productService.AddProduct(&product)
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(product)
	return nil
}
