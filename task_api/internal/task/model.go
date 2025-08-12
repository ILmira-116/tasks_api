package task

import (
	"time"
)

type Status string

const (
	Waiting  Status = "waiting"
	Active   Status = "active"
	Finished Status = "finished"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
