package jobs

import (
	"github.com/hibiken/asynq"
	"log"
	"pheasant-api/app/helper"
)

var queueClient *asynq.Client = nil

func getQueueClient() *asynq.Client {
	if queueClient != nil {
		return queueClient
	}

	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     helper.GetEnv("REDIS_HOST", ""),
		Username: helper.GetEnv("REDIS_USERNAME", ""),
		Password: helper.GetEnv("REDIS_PASSWORD", ""),
	})
}

func Enqueue(task *asynq.Task, options ...asynq.Option) error {
	client := getQueueClient()

	info, err := client.Enqueue(task, options...)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
		return err
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	return nil
}
