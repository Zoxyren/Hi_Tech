package repository

import (
	"Hi_Tech/internal/model"
	"database/sql"
	"fmt"
	"log/slog"
)

type UserRepository struct {
	user *model.User
	db   *sql.DB
}

func (u *UserRepository) CreateUserAccount(users *model.User) (*model.User, error) {
	// Ensure Cart ID is created first
	cartID, err := InsertIntoCarts(u.db, users.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into carts: %w", err)
	}
	users.CartID = cartID

	stmt, err := u.db.Prepare("INSERT INTO users(username, password, email, cart_id) VALUES($1, $2, $3, $4)")
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	defer func() {
		err := stmt.Close()
		if err != nil {
			slog.Error("Failed to close statement", "error")
		}
	}()

	_, err = stmt.Exec(users.Username, users.Password, users.Email, users.CartID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	return users, nil
}

func InsertIntoCarts(db *sql.DB, userID int) (int, error) {
	// Insert a new cart record and return the cart ID
	result, err := db.Exec("INSERT INTO carts (user_id) VALUES ($1)", userID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert into carts: %w", err)
	}

	cartID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert id: %w", err)
	}

	return int(cartID), nil
}
