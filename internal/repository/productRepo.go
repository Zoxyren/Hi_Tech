package repository

import (
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID            int     `db:"product_id"`
	Name          string  `db:"name"`
	ImageURL      string  `db:"image_url"`
	Description   string  `db:"description"`
	Price         float64 `db:"price"`
	StockQuantity int     `db:"stock_quantity"`
}

type ProductRepository struct {
	db *sqlx.DB
}

func (r *ProductRepository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Select(&products, "SELECT product_id, name, image_url, description, price, stock_quantity FROM products")
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Add(product Product) error {
	query := `INSERT INTO products (name, image_url, description, price, stock_quantity) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, product.Name, product.ImageURL, product.Description, product.Price, product.StockQuantity)
	return err
}

func (r *ProductRepository) Remove(id int) error {
	query := `DELETE FROM cart_items WHERE product_id = $1; DELETE FROM products WHERE product_id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
