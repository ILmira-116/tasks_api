# Task API

API для управления задачами с фильтрацией по статусу.

## Функционал

- **GET** `/tasks` — получить список задач (опционально фильтр по статусу: `waiting`, `active`, `finished`)
- **GET** `/tasks/{id}` — получить задачу по ID
- **POST** `/tasks` — создать новую задачу

## Сборка и запуск

1. Клонировать репозиторий:
   ```bash
   git clone git@github.com:ILmira-116/tasks_api.git
   cd tasks_api

2. Запустить приложение:
go run cmd/main.go  

## API будет доступно на:
http://localhost:8080

## Пример запросов
# Получить все задачи:
GET http://localhost:8080/tasks

# Получить задачи со статусом:
GET http://localhost:8080/tasks?status=active

# Получить задачу по ID:
GET http://localhost:8080/tasks/{id}

# Создать задачу:
POST http://localhost:8080/tasks
Content-Type: application/json

{
  "ID": "1",
  "Title": "Новая задача",
  "Description": "Описание задачи",
  "Status": "active"
}

# Затраченное время: 
Около 4 часов (включая настройку окружения, разработку API, тестирование в Incomnia и загрузку в GitHub).
