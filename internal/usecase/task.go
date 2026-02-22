package usecase

import (
	"practice3/internal/models"
	"practice3/internal/repository"
)

type TaskUseCase struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: r}
}


func (u *TaskUseCase) GetAll() ([]models.Task, error) {
	return u.repo.GetAll()
}

func (u *TaskUseCase) GetByID(id string) (*models.Task, error) {
	return u.repo.GetByID(id)
}

func (u *TaskUseCase) Create(task *models.Task) error {
	return u.repo.Create(task)
}

func (u *TaskUseCase) Update(task *models.Task) error {
	return u.repo.Update(task)
}	

func (u *TaskUseCase) Delete(id string) error {
	return u.repo.Delete(id)
}