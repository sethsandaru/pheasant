package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

type MailRequest struct {
	from    string
	to      []string
	subject string
	body    string

	// smtp configuration
	smtpHost     string
	smtpUsername string
	smtpPassword string
	smtpPort     int
}

const (
	MIME                  = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	BaseEmailTemplatePath = "./resources/templates/emails/"
)

func NewMailRequest(to []string, subject string) *MailRequest {
	return &MailRequest{
		to:      to,
		subject: subject,

		smtpHost:     GetEnv("SMTP_HOST", ""),
		smtpUsername: GetEnv("SMTP_USERNAME", ""),
		smtpPassword: GetEnv("SMTP_PASSWORD", ""),
		smtpPort:     GetIntEnv("SMTP_PORT", 25),
	}
}

func (mailRequest *MailRequest) parseTemplate(fileName string, data interface{}) error {
	templateFile, err := template.ParseFiles(BaseEmailTemplatePath + fileName)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	if err = templateFile.Execute(buffer, data); err != nil {
		return err
	}

	mailRequest.body = buffer.String()

	return nil
}

func (mailRequest *MailRequest) sendMail() bool {
	body := "To: " + mailRequest.to[0] + "\r\nSubject: " + mailRequest.subject + "\r\n" + MIME + "\r\n" + mailRequest.body
	SMTP := fmt.Sprintf("%s:%d", mailRequest.smtpHost, mailRequest.smtpPort)

	err := smtp.SendMail(
		SMTP,
		smtp.PlainAuth(
			"",
			mailRequest.smtpUsername,
			mailRequest.smtpPassword,
			mailRequest.smtpHost,
		),
		mailRequest.smtpUsername,
		mailRequest.to,
		[]byte(body),
	)
	if err != nil {
		return false
	}

	return true
}

func (mailRequest *MailRequest) Send(templateName string, items interface{}) {
	err := mailRequest.parseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}

	if ok := mailRequest.sendMail(); ok {
		log.Printf("Email has been sent to %s\n", mailRequest.to)
	} else {
		log.Printf("Failed to send the email to %s\n", mailRequest.to)
	}
}
