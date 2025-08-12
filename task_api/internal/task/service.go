package task

import (
	"errors"
	"task_api/task_api/internal/logger"
	"time"
)

type TaskService struct {
	repo   *Repository
	logger *logger.Logger
}

func NewTaskService(repo *Repository, logger *logger.Logger) *TaskService {
	return &TaskService{
		repo:   repo,
		logger: logger,
	}
}

func (s *TaskService) CreateTask(t *Task) error {
	if t.ID == "" {
		return errors.New("task ID is required")
	}

	t.CreatedAt = time.Now()
	s.repo.Create(t)
	s.logger.Info("Created task with ID " + t.ID)
	return nil
}

func (s *TaskService) GetTaskById(id string) (*Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) GetTasksAll(status *Status) ([]*Task, error) {
	return s.repo.GetByAll(status)
}
