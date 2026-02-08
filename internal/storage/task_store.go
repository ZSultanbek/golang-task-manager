package storage

import (
	"sync"

	"assignment1/internal/models"
)

type TaskStore struct {
	mu     sync.RWMutex
	tasks  map[int]models.Task
	nextID int
}
