package repository

import (
	"database/sql"
	"log"
)

type CartRepository struct {
	db *sql.DB
}

func (cr *CartRepository) GetCartIdByUserId(userID int) int {
	var cartID int
	err := cr.db.QueryRow("SELECT cart_id FROM carts WHERE user_id = $1;", userID).Scan(&cartID)
	if err == sql.ErrNoRows {
		log.Fatal("Cart not found!")
	} else if err != nil {
		log.Fatal(err)
	}

	return cartID
}

func (cr *CartRepository) CreateCart(userID int) {
	_, err := cr.db.Exec("INSERT INTO carts (user_id) VALUES ($1);", userID)
	if err != nil {
		log.Fatal(err)
	}
}
