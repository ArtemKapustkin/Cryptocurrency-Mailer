package mailer

import (
	"net/smtp"
	"os"
)

type Mailer struct {
	server   string
	port     string
	email    string
	password string
}

func NewMailer(server, port string) *Mailer {
	return &Mailer{
		server:   server,
		port:     port,
		email:    os.Getenv("SENDER_EMAIL"),
		password: os.Getenv("SENDER_PASSWORD"),
	}
}

func (m *Mailer) SendEmail(email, message string) error {
	auth := smtp.PlainAuth("", m.email, m.password, m.server)

	err := smtp.SendMail(m.server+":"+m.port, auth, m.email, []string{email}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
