package main

import (
	"net/http"

	"task_api/task_api/internal/config"
	"task_api/task_api/internal/graceful"
	"task_api/task_api/internal/logger"
	"task_api/task_api/internal/task"
)

func main() {
	//1. Загрузка конфига
	cfg := config.Load()

	//2. Инициализация логгера
	lg, err := logger.NewLogger(cfg.LogFilePath, 100)
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	defer lg.Close()
	lg.Info("Logger initialized")

	//3. Инициализация репозитория
	repo := task.NewRepository(lg)
	lg.Info("Repository initialized")

	//4. Инициализация сервиса
	service := task.NewTaskService(repo, lg)
	lg.Info("Service initialized")

	//5. Создаем хендлер
	handler := task.NewTaskHandler(service)
	lg.Info("Handler initialized")

	//6. Роутер
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetTasksAll(w, r)
		case http.MethodPost:
			handler.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/tasks/", handler.GetTaskById)

	//7. Сервер
	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: mux,
	}

	//8. Запуск сервера в отдельной горутине
	go func() {
		lg.Info("Server starting on :" + cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil {
			lg.Error("Server error: " + err.Error())
		}
	}()

	// 9. Ожидание и graceful shutdown
	graceful.ShutdownServer(server, cfg.ShutdownTimeout)
	lg.Info("Server stopped gracefully")

}
