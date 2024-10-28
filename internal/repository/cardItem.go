package repository

import (
	"database/sql"
	"log"
)

type CartItem struct {
	CartItemID int
	Name       string
	ImageURL   string
	Price      float64
	Quantity   int
}

type CartRecord struct {
	CartID    int
	ProductID int
	Quantity  int
}

type CartItemsRepository struct {
	db *sql.DB
}

func NewCartItemsRepository(db *sql.DB) *CartItemsRepository {
	return &CartItemsRepository{db: db}
}

func (cir *CartItemsRepository) GetItemsByCartId(cartID int) []CartItem {
	rows, err := cir.db.Query("SELECT cart_items.cart_item_id, products.name, products.image_url, products.price, cart_items.quantity FROM products INNER JOIN cart_items ON products.product_id = cart_items.product_id WHERE cart_id = $1;", cartID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cartItems []CartItem
	for rows.Next() {
		var cartItem CartItem
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Name, &cartItem.ImageURL, &cartItem.Price, &cartItem.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems
}

func (cir *CartItemsRepository) UpdateItemQuantity(cartItemID int, quantity int) {
	_, err := cir.db.Exec("UPDATE cart_items SET quantity = $1 WHERE cart_item_id = $2;", quantity, cartItemID)
	if err != nil {
		log.Fatal(err)
	}
}

func (cir *CartItemsRepository) AddItemToCart(cartRecord CartRecord) {
	_, err := cir.db.Exec("INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ($1, $2, $3);", cartRecord.CartID, cartRecord.ProductID, cartRecord.Quantity)
	if err != nil {
		log.Fatal(err)
	}
}
