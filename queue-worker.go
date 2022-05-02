package main

import (
	"github.com/hibiken/asynq"
	"log"
	"pheasant-api/app/helper"
	"pheasant-api/app/jobs"
)

func main() {
	workerServer := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     helper.GetEnv("REDIS_HOST", ""),
			Username: helper.GetEnv("REDIS_USERNAME", ""),
			Password: helper.GetEnv("REDIS_PASSWORD", ""),
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: helper.GetIntEnv("QUEUE_WORKER_NUMBER", 10),
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(jobs.InitForgotPasswordJob().GetName(), jobs.InitForgotPasswordJob().Handle)

	if err := workerServer.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
