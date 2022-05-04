package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"pheasant-api/app/helper"
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

const ForgotPasswordTemplate = "forgot_password_template.html"

type forgotPasswordTemplateData struct {
	user             models.User
	resetPasswordUrl string
}

func InitForgotPasswordJob() ForgotPasswordJob {
	return &forgotPasswordDependencies{
		name: "email:forgot-password",
	}
}

func (job *forgotPasswordDependencies) GetName() string {
	return job.name
}

// Dispatch will dispatch the queue task
func (job *forgotPasswordDependencies) Dispatch(user models.User) error {
	payload, err := json.Marshal(user)
	if err != nil {
		return err
	}

	task := asynq.NewTask(job.name, payload)
	return Enqueue(task)
}

// Handle will run the task and send out email to user
func (job *forgotPasswordDependencies) Handle(ctx context.Context, t *asynq.Task) error {
	var user models.User
	if err := json.Unmarshal(t.Payload(), &user); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	// create forgot password token
	resetPasswordToken := helper.GenerateUUID()
	resetPasswordUrl := helper.GetEnv("FRONTEND_FORGOT_PASSWORD_URL", "http://localhost/forgot-password/") + resetPasswordToken

	// send email
	mailRequest := helper.NewMailRequest([]string{user.Email}, "Forgot Password Instruction")
	mailRequest.Send(ForgotPasswordTemplate, forgotPasswordTemplateData{
		user:             user,
		resetPasswordUrl: resetPasswordUrl,
	})

	return nil
}
