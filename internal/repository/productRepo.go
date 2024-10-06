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

func (r *ProductRepository) FetchProductById(id int) (*model.Products, error) {
	var product model.Products
	row := r.Db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id)
	err := row.Scan(&product.ProductID, &product.Name, &product.Price)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &product, err
}
