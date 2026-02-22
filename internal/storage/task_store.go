package storage

import (
	"sync"

	"practice3/internal/models"
)

type TaskStore struct {
	mu     sync.RWMutex
	tasks  map[int]models.Task
	nextID int
}
