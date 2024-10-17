package model

type User struct {
	CartID   int    `json:"cartid"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
