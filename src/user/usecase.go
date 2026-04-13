package user

import "advanced/structs"

type userUsecase struct {
	repo IUserRepository
}

func ProvideUserUsecase(repo IUserRepository) IUserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetAllUsers(limit int, page int) ([]structs.User, error) {
	offset := (page - 1) * limit
	return u.repo.FindAll(limit, offset)
}

func (u *userUsecase) GetUser(id int, name, email string) (*structs.User, error) {
	return u.repo.Find(id, name, email)
}

func (u *userUsecase) RegisterUser(user *structs.User) error {
	return u.repo.Create(user)
}
