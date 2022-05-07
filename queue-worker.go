package main

import (
	"github.com/hibiken/asynq"
	"log"
	"pheasant-api/app/helper"
	"pheasant-api/app/jobs"
)

func main() {
	redisHost := helper.GetEnv("REDIS_HOST", "")
	redisUser := helper.GetEnv("REDIS_USERNAME", "")
	redisPass := helper.GetEnv("REDIS_PASSWORD", "")
	if redisHost == "" || redisUser == "" || redisPass == "" {
		panic("Missing Redis configuration. Aborted")
	}

	workerServer := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     redisHost,
			Username: redisUser,
			Password: redisPass,
		},
		asynq.Config{
			Concurrency: helper.GetIntEnv("QUEUE_WORKER_NUMBER", 10),
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	forgotPasswordJob := jobs.InitForgotPasswordJob()
	mux.HandleFunc(forgotPasswordJob.GetName(), forgotPasswordJob.Handle)

	if err := workerServer.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
