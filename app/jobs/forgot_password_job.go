package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"pheasant-api/app/models"
)

type ForgotPasswordJob interface {
	GetName() string

	Dispatch(user models.User) error

	Handle(ctx context.Context, t *asynq.Task) error
}

type forgotPasswordDependencies struct {
	name string
}

func InitForgotPasswordJob() ForgotPasswordJob {
	return &forgotPasswordDependencies{
		name: "email:forgot-password",
	}
}

func (job *forgotPasswordDependencies) GetName() string {
	return job.name
}

func (job *forgotPasswordDependencies) Dispatch(user models.User) error {
	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}

	task := asynq.NewTask(job.name, payload)
	return Enqueue(task)
}

func (job *forgotPasswordDependencies) Handle(ctx context.Context, t *asynq.Task) error {
	var user models.User
	if err := json.Unmarshal(t.Payload(), &user); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	// create forgot password token

	// send email

	return nil
}
