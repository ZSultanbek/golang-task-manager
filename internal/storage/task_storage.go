package storage

import (
	"task-manager/internal/models"
)

func (s *TaskStore) GetAllTasks() ([]models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *TaskStore) GetTaskByID(id int) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

func (s *TaskStore) FilterTasks(done bool) ([]models.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]models.Task, 0)
	for _, task := range s.tasks {
		if task.Done == done {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func (s *TaskStore) CreateTask(taskTitle string) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := models.Task{
		ID:    s.nextID,
		Title: taskTitle,
		Done:  false,
	}

	s.tasks[task.ID] = task
	s.nextID++

	return task, nil
}

func (s *TaskStore) UpdateTask(id int, done bool) (bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return  false
	}

	task.Done = done
	s.tasks[id] = task

	return true
}

func (s *TaskStore) DeleteTask(id int) (bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.tasks[id]
	if !exists {
		return false
	}

	delete(s.tasks, id)
	return true
}

var GlobalTaskStore = &TaskStore{
	tasks:  make(map[int]models.Task),
	nextID: 1,
}