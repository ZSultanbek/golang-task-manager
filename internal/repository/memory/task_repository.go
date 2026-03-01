package memory

import (
	"errors"
	"task-manager/internal/models"
)

type TaskRepo struct {
	tasks  []models.Task
	nextID int
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		tasks:  []models.Task{},
		nextID: 1,
	}
}




func (r *TaskRepo) GetAll() ([]models.Task, error) {
	return r.tasks, nil
}

func (r *TaskRepo) GetByID(id int) (*models.Task, error) {
	for _, t := range r.tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *TaskRepo) Create(task *models.Task) (int, error) {
	task.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, *task)
	return task.ID, nil
}

func (r *TaskRepo) Update(task *models.Task) error {
	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = *task
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *TaskRepo) Delete(id int) error {
	for i, t := range r.tasks {
		if t.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}