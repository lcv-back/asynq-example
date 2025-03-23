A minimal, extensible Go project demonstrating background task processing with Asynq.

## Project Structure

```
ASYNC-EXAMPLE/
├── cmd/
│   ├── client/main.go      # HTTP API enqueuing tasks
│   └── worker/main.go      # Asynq worker processing tasks
├── internal/
│   ├── handler/handler.go  # Task handlers
│   └── tasks/tasks.go      # Task definitions & constructors
├── pkg/
│   └── redis/redis.go      # Redis connection config
├── docs/structure.md       # Project structure overview
├── go.mod
├── go.sum
└── README.md
```
