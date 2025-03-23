package main

//Ứng dụng worker để xử lý tasks
import (
	"log"

	"asynq-example/internal/handler"
	"asynq-example/internal/tasks"
	"asynq-example/pkg/redis"

	"github.com/hibiken/asynq"
)

func main() {
	server := asynq.NewServer(redis.RedisOpt(), asynq.Config{Concurrency: 10})
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailDelivery, handler.HandleEmailDelivery)

	if err := server.Run(mux); err != nil {
		log.Fatalf("worker failed: %v", err)
	}
}
