package controller

import (
	"Hi_Tech/internal/model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	products *[]model.Products
	db       *sql.DB
}

func (p *Product) GetAllProducts(w http.ResponseWriter, r *http.Request) []model.Products {
	w.Header().Set("Content-Type", "application/json")
	rows, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		log.Printf("error fetching products from database. Err: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	defer rows.Close()

	var products []model.Products
	for rows.Next() {
		var product model.Products
		err := rows.Scan(&product.ProductID, &product.Name, &product.Price)
		if err != nil {
			log.Printf("error scanning row. Err: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
		products = append(products, product)
	}

	jsonResp, err := json.Marshal(products)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
	return products
}
