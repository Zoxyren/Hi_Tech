package repository

import (
	"Hi_Tech/internal/errorHandling"
	"Hi_Tech/internal/model"
	"database/sql"
	"errors"
	"log/slog"
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
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return &product, err
}

func (r *ProductRepository) InsertItem(product *model.Products) error {
	stmt, err := r.Db.Prepare("INSERT INTO products(name, price, image, description, stock) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return errorHandling.ErrorAddingProduct
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			slog.Error("Failed to insert Item", "error", errorHandling.ErrInsertingStatement)
		}
	}(stmt)

	_, err = stmt.Exec(product.Name, product.Price, product.Image, product.Description, product.Stock)
	return err
}
