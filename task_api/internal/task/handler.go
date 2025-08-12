package task

import (
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	service *TaskService
}

func NewTaskHandler(service *TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var t Task

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTask(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)

}

func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("/tasks/"):]

	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTaskById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func (h *TaskHandler) GetTasksAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	statusParam := r.URL.Query().Get("status")

	var status *Status

	if statusParam != "" {
		s := Status(statusParam)
		switch s {
		case Waiting, Active, Finished:
			status = &s
		default:
			http.Error(w, "Invalid status value", http.StatusBadRequest)
			return
		}
	}

	tasks, err := h.service.GetTasksAll(status)
	if err != nil {
		http.Error(w, "Failed to get tasks: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}
