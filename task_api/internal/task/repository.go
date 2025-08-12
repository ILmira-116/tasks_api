package task

import (
	"errors"
	"fmt"
	"sync"

	"task_api/task_api/internal/logger"
)

type Repository struct {
	mu     sync.RWMutex
	tasks  map[string]*Task
	logger *logger.Logger
}

func NewRepository(logger *logger.Logger) *Repository {
	return &Repository{
		tasks:  make(map[string]*Task),
		logger: logger,
	}
}

func (r *Repository) Create(task *Task) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[task.ID] = task
	r.logger.Info(fmt.Sprintf("Created task with ID %s", task.ID))
}

func (r *Repository) GetByID(id string) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, ok := r.tasks[id]
	if !ok {
		errMsg := fmt.Sprintf("Task with ID %s not found", id)
		r.logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	r.logger.Info(fmt.Sprintf("Fetched task with ID %s", id))

	return task, nil

}

func (r *Repository) GetByAll(statusFilter *Status) ([]*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []*Task

	for _, task := range r.tasks {
		if statusFilter == nil || task.Status == *statusFilter {
			list = append(list, task)
			r.logger.Info(fmt.Sprintf("Task found: ID=%s, Status=%s", task.ID, task.Status))
		}
	}

	return list, nil

}
