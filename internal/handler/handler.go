package handler

//Xử lý logic của các tasks
import (
	"context"
	"encoding/json"
	"log"

	"asynq-example/internal/tasks"

	"github.com/hibiken/asynq"
)

func HandleEmailDelivery(ctx context.Context, t *asynq.Task) error {
	var p tasks.EmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf("Sending email to=%s subject=%s body=%s", p.To, p.Subject, p.Body)
	return nil
}
