package tasks

//Định nghĩa các loại task và payload

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypeEmailDelivery = "email:deliver"

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func NewEmailTask(to, subject, body string) (*asynq.Task, error) {
	payload := EmailPayload{To: to, Subject: subject, Body: body}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, b), nil
}
