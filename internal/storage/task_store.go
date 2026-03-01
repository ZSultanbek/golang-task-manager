package storage

import (
	"sync"

	"task-manager/internal/models"
)

type TaskStore struct {
	mu     sync.RWMutex
	tasks  map[int]models.Task
	nextID int
}
