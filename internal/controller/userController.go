package controller

import (
	"Hi_Tech/internal/model"
	"Hi_Tech/internal/repository"
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	user model.User
	db   *sql.DB
	model.User
}

func (u *User) RegisterUser(w http.ResponseWriter, r *http.Request) (error, error) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err, nil
	}

	// Assuming there is a method called Register in the model.User struct
	userRepository := repository.UserRepository{}
	newUser, err := userRepository.Register(user) // Fixed: removed unnecessary type conversion
	if err != nil {
		return err, nil
	}

	user = newUser
	json.NewEncoder(w).Encode(user)
	return nil, nil
}
