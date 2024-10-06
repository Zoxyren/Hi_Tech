package repository

import (
	"Hi_Tech/internal/model"
	"database/sql"
)

type ProductRepository struct {
	Db *sql.DB
}

func (r *ProductRepository) FetchAll() ([]model.Products, error) {
	rows, err := r.Db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Products
	for rows.Next() {
		var product model.Products
		err := rows.Scan(&product.ProductID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
