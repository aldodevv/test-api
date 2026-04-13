package user

import (
	"advanced/structs"
)

type IUserRepository interface {
	FindAll(limit int, offset int) ([]structs.User, error)
	Find(id int, name, email string) (*structs.User, error)
	Create(user *structs.User) error
}

type IUserUsecase interface {
	GetAllUsers(limit int, page int) ([]structs.User, error)
	GetUser(id int, name, email string) (*structs.User, error)
	RegisterUser(user *structs.User) error
}
