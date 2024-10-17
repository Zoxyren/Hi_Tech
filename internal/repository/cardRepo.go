package repository

import (
	"Hi_Tech/internal/model"
	"database/sql"
)

type User struct {
	user *model.User
	db   *sql.DB
}
