# Task Management App (Go)

Small task-queue and worker example written in Go. Provides an HTTP API to create and list tasks and an in-memory priority queue processed by workers.

## Features
- Add tasks via HTTP POST `/tasks`
- List tasks via HTTP GET `/tasks`
- Priority-based in-memory queue and worker processing
- Simple logger integration

## Requirements
- Go 1.25+ (modules enabled)

## Build & Run
Build:
```sh
go build ./cmd/task_management_app
./task_management_app
```
Run directly:
```sh
go run ./cmd/task_management_app
```

## HTTP API
- POST /tasks — add a new task (JSON body)
- GET /tasks — list all tasks

Example (curl):
```sh
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"id":"task1","priority":1,"payload":"do work"}'

curl http://localhost:8080/tasks
```

## Tests
Run all tests:
```sh
go test ./...
```

## Key files
- cmd/taskmanager/main.go — app entrypoint
- api/handler/task_handlers.go — HTTP handlers
- api/handler/request.go — request DTOs / validation
- internal/router/routers.go — router setup
- core/app/app.go — application task holder
- internal/tasks/task.go — task model
- internal/tasks/queue.go — queue implementation
- internal/tasks/worker.go — worker loop
- internal/logging/logger.go — logger
- tests/api_test.go, tests/queue_test.go, tests/task_test.go — tests
