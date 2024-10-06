package model

import (
	"github.com/google/uuid"
)

type image struct {
	Url string `json:"url"`
}
type Cart struct {
	CartID int64 `json:"cart_id"`
	UserID uuid.UUID
}

type Products struct {
	ProductID   int    `json:"product_id"`
	Name        string `json:"name"`
	Image       image  `json:"image"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type CartItems struct {
	CartItemID int64 `json:"cart_item_id"`
	CartID     int64 `json:"cart_id"`
	ProductID  int64 `json:"product_id"`
	Quantity   int   `json:"quantity"`
}
