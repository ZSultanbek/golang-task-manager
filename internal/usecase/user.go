package usecase

import (
	"practice3/internal/models"
	"practice3/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) GetAll() ([]models.User, error) {
	return u.repo.GetAll()
}

func (u *UserUsecase) GetByID(id int) (*models.User, error) {
	return u.repo.GetByID(id)
}

func (u *UserUsecase) Create(user *models.User) (int, error) {
	return u.repo.Create(user)
}

func (u *UserUsecase) Update(user *models.User) error {
	return u.repo.Update(user)
}

func (u *UserUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}