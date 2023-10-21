package main

import (
	"async-quickstart/task"
	"github.com/hibiken/asynq"
	"log"
)

// workers.go
func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

type emailTaskPayload struct {
	// ID for the email recipient.
	UserID int
}
