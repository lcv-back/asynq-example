# asynq-example

```markdown
# Asynq Example

## Setup

1. Run Redis on default port 6379.
2. Clone repo and run `go mod tidy`.

### Start API Client
```

go run cmd/client/main.go

```

### Start Worker
```

go run cmd/worker/main.go

```

### Enqueue Task
```

curl -X POST http://localhost:8080/enqueue/email -d '{"to":"test@example.com","subject":"Hi","body":"Hello from Asynq"}' -H "Content-Type: application/json"

```

```
