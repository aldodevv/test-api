package user

import (
	"advanced/structs"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func ProvideUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll(limit int, offset int) ([]structs.User, error) {
	var users []structs.User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *userRepository) Find(id int, name, email string) (*structs.User, error) {
	var user structs.User
	// GORM Where mendeteksi properti non-zero (jika id=0, dihiraukan)
	err := r.db.Where(&structs.User{
		Model: gorm.Model{ID: uint(id)},
		Name:  name,
		Email: email,
	}).First(&user).Error
	return &user, err
}

func (r *userRepository) Create(user *structs.User) error {
	return r.db.Create(user).Error
}
