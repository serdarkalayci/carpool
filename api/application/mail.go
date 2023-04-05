// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"errors"
	"os"
	"strconv"

	"github.com/go-mail/mail"
)

func sendEmail(to, subject, body string) error {
	smtpServer := os.Getenv("SmtpServer")
	smtpPortS := os.Getenv("SmtpPort")
	smtpPort := 0
	if smtpPortS != "" {
		smtpPort, _ = strconv.Atoi(smtpPortS)
	}
	smtpUser := os.Getenv("SmtpUser")
	smtpPassword := os.Getenv("SmtpPassword")

	if smtpServer == "" || smtpPort == 0 || smtpUser == "" || smtpPassword == "" {
		return errors.New("SMTP server not configured")
	}

	m := mail.NewMessage()

	m.SetHeader("From", "info@bidunyaoy.com")

	m.SetHeader("To", to)

	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)

	d := mail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPassword)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	if err := d.DialAndSend(m); err != nil {

		panic(err)
	}
	return nil
}
