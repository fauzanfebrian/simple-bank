package mail

import (
	"fmt"
	"net/smtp"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/jordan-wright/email"
)

type EmailSender struct {
	name             string
	fromEmailAddress string
	password         string
	username         string
	host             string
	port             string
}

func NewEmailSender(config util.Config) *EmailSender {
	return &EmailSender{
		name:             config.EmailSenderName,
		fromEmailAddress: config.EmailSenderAddress,
		password:         config.EmailSenderPassword,
		username:         config.EmailSenderUsername,
		host:             config.EmailSenderHost,
		port:             config.EmailSenderPort,
	}
}

func (sender *EmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.HTML = []byte(content)

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpServerAddress := fmt.Sprintf("%s:%s", sender.host, sender.port)

	smtpAuth := smtp.PlainAuth("", sender.username, sender.password, sender.host)
	return e.Send(smtpServerAddress, smtpAuth)
}
