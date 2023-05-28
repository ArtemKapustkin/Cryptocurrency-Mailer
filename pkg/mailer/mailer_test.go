package mailer

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestSendingEmail(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Errorf("error loading .env file")
	}

	mailer := NewMailer("smtp.gmail.com", "587")

	err = mailer.SendEmail("a.kapustkin.2003@gmail.com", "test_message")
	if err != nil {
		t.Errorf("failure occurs while sending email: %v", err)
	}
}
