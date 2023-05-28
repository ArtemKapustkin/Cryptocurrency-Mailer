package service

import (
	"bitcoin-exchange-rate/internal/model"
	"fmt"
	"log"
)

type SubscriberRepository interface {
	GetAll() ([]*model.Subscriber, error)
}

type Mailer interface {
	SendEmail(email, value string) error
}

type MailerService struct {
	subscriberRepository SubscriberRepository
	mailer               Mailer

	messageToSend string
}

func NewMailerService(subscriberRepository SubscriberRepository, mailer Mailer) *MailerService {
	return &MailerService{
		subscriberRepository: subscriberRepository,
		mailer:               mailer,
		messageToSend:        "Subject: BTCUAH Exchange Rate Update\n\nDear subscriber,\n\nHere is current BTCUAH exchange rate: %s\n\nSincerely,\nArtem Kapustkin Mailer",
	}
}

func (s *MailerService) SendValueToAllEmails(value string) error {
	subscribers, err := s.subscriberRepository.GetAll()
	if err != nil {
		return err
	}

	for _, subscriber := range subscribers {
		message := fmt.Sprintf(s.messageToSend, value)

		err := s.mailer.SendEmail(subscriber.GetEmail(), message)
		if err != nil {
			log.Printf("Failed to send email to %s: %s", subscriber.GetEmail(), err)
			return err
		}

		log.Printf("Email sent successfully to %s", subscriber.GetEmail())
	}

	return nil
}
