package client

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

type EmailTaskPayLoad struct {
	UserID int
}

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	/// create a task with typeName and payload
	payload, err := json.Marshal(EmailTaskPayLoad{UserID: 42})
	if err != nil {
		log.Fatal(err)
	}

	task1 := asynq.NewTask("email:welcome", payload)
	task2 := asynq.NewTask("email:reminder", payload)

	/// process the task immediately
	info, err := client.Enqueue(task1)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf(" [*] Successfully enqueued task: %+v", info)
	}

	/// process the task 60 seconds later
	info, err = client.Enqueue(task2, asynq.ProcessIn(60*time.Second))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf(" [*] Successfully enqueued task: %+v", info)
	}

}
