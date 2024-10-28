package controller

import (
	"Hi_Tech/internal/model"
	"Hi_Tech/internal/repository"
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	user *model.User
	db   *sql.DB
}

func (u *User) RegisterUser(w http.ResponseWriter, r *http.Request) (error, error) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err, nil
	}

	userRepository := repository.UserRepository{}
	newUser, err := userRepository.Register(user)
	if err != nil {
		// Log the error for debugging
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return err, nil
	}

	json.NewEncoder(w).Encode(newUser)
	return nil, nil
}
