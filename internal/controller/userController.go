package controller

import (
	"Hi_Tech/internal/errorHandling"
	"Hi_Tech/internal/model"
	"Hi_Tech/internal/repository"
	"Hi_Tech/internal/services"
	_ "Hi_Tech/internal/services"
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	user *model.User
	db   *sql.DB
}

func (u *User) RegisterUser(w http.ResponseWriter, r *http.Request) error {
	// Validate request body
	var user *model.User
	productService := services.UserService{Repo: repository.UserRepository{}, Db: u.db}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return errorHandling.ErrCreatingUser
	}
	insert, err := productService.CreateUser(user)
	if err != nil {
		return err
	}
	err = json.NewEncoder(w).Encode(insert)
	if err != nil {
		return err
	}
	return nil
}
