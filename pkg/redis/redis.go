package redis

//Code liên quan đến Redis (nếu tách riêng)
import (
	"os"

	"github.com/hibiken/asynq"
)

func RedisOpt() asynq.RedisClientOpt {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	return asynq.RedisClientOpt{Addr: addr}
}
