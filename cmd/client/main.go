// Ứng dụng client để enqueue tasks
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"asynq-example/internal/tasks"
	"asynq-example/pkg/redis"

	"github.com/hibiken/asynq"
)

func main() {
	client := asynq.NewClient(redis.RedisOpt())
	defer client.Close()

	http.HandleFunc("/enqueue/email", func(w http.ResponseWriter, r *http.Request) {
		var payload tasks.EmailPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task, err := tasks.NewEmailTask(payload.To, payload.Subject, payload.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		info, err := client.Enqueue(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(info)
	})

	log.Println("Client API running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
