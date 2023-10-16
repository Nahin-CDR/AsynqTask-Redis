package workers

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

func main() {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}

}
