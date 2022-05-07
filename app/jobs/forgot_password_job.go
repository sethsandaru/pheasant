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
	User             models.User
	ResetPasswordUrl string
}

type forgotPasswordJobPayload struct {
	User   models.User
	UserID uint64
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
	payload, err := json.Marshal(&forgotPasswordJobPayload{
		User:   user,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	task := asynq.NewTask(job.name, payload)
	return Enqueue(task)
}

// Handle will run the task and send out email to user
func (job *forgotPasswordDependencies) Handle(ctx context.Context, t *asynq.Task) error {
	var payload forgotPasswordJobPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v", err)
	}

	// create forgot password token
	payload.User.ID = payload.UserID
	resetPasswordToken, err := models.GetForgotPasswordTokenModel().CreateForUser(payload.User)
	if err != nil {
		return fmt.Errorf("Failed to create token for User")
	}

	// generate URL
	resetPasswordUrl := helper.GetEnv("FRONTEND_FORGOT_PASSWORD_URL", "http://localhost/forgot-password/") + resetPasswordToken.Token

	// send email
	mailRequest := helper.NewMailRequest([]string{payload.User.Email}, "Forgot Password Instruction")
	mailRequest.Send(ForgotPasswordTemplate, forgotPasswordTemplateData{
		User:             payload.User,
		ResetPasswordUrl: resetPasswordUrl,
	})

	return nil
}
