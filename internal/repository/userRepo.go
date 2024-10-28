package repository

import (
	"Hi_Tech/internal/model"
	"database/sql"
	"errors"
	"log"
)

type User struct {
	*model.User
}

type Credentials struct {
	username string
	password string
}

type UserRepository struct {
	db             *sql.DB
	cartRepository *CartRepositorys
}

func (ur *UserRepository) Login(credentials Credentials) bool {
	user := &User{}
	err := ur.db.QueryRow("SELECT * FROM users WHERE username = $1 AND password = $2;", credentials.username, credentials.password).Scan(&user.UserID, &user.Username, &user.Email, &user.Password)
	if errors.Is(err, sql.ErrNoRows) {
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	user.CartID = ur.cartRepository.GetCartIdByUserId(user.UserID)
	return true
}

func (ur *UserRepository) Register(user model.User) (int, error) {
	err := ur.db.QueryRow("INSERT INTO users(username, email, password) VALUES ($1, $2, $3) RETURNING user_id;", user.Username, user.Email, user.Password).Scan(&user.UserID)
	if err != nil {
		log.Fatal(err)
	}

	ur.cartRepository.CreateCart(user.UserID)
	return user.UserID, nil
}

func (ur *UserRepository) Delete(user User) int {
	err := ur.db.QueryRow("SELECT * FROM users WHERE user_id = $1 AND password = $2;", user.UserID, user.Password).Scan()
	if errors.Is(err, sql.ErrNoRows) {
		log.Fatal("Wrong password!")
	} else if err != nil {
		log.Fatal(err)
	}

	cartID := ur.cartRepository.GetCartIdByUserId(user.UserID)
	_, err = ur.db.Exec("DELETE FROM cart_items WHERE cart_id = $1;", cartID)
	if err != nil {
		log.Fatal(err)
	}

	res, err := ur.db.Exec("DELETE FROM carts WHERE user_id = $1; DELETE FROM users WHERE user_id = $1;", user.UserID)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return int(rowsAffected)
}

type CartRepositorys struct {
	db *sql.DB
}

func NewCartsRepository(db *sql.DB) *CartRepositorys {
	return &CartRepositorys{db: db}
}

func (cr *CartRepositorys) GetCartIdByUserId(userID int) int {
	var cartID int
	err := cr.db.QueryRow("SELECT cart_id FROM carts WHERE user_id = $1;", userID).Scan(&cartID)
	if err == sql.ErrNoRows {
		log.Fatal("Cart not found!")
	} else if err != nil {
		log.Fatal(err)
	}

	return cartID
}

func (cr *CartRepositorys) CreateCart(userID int) {
	_, err := cr.db.Exec("INSERT INTO carts (user_id) VALUES ($1);", userID)
	if err != nil {
		log.Fatal(err)
	}
}
