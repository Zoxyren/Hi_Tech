package services

import (
	"Hi_Tech/internal/model"
	"Hi_Tech/internal/repository"
	"database/sql"
)

type UserService struct {
	user *model.User
	Db   *sql.DB
	Repo repository.UserRepository
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	return u.Repo.CreateUserAccount(user)
}
