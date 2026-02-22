package repository

import (
	"practice3/internal/models"
	"practice3/internal/repository/postgres"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user *models.User) (int, error)
	Update(user *models.User) error
	Delete(id int) error
}

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id string) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
}

type Repositories struct {
	UserRepo UserRepository
}

func NewRepositories(db *postgres.Dialect) *Repositories {
	return &Repositories{
		UserRepo: postgres.NewUserRepository(db),
	}
}