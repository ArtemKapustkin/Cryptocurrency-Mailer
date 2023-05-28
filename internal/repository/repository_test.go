package repository

import (
	"bitcoin-exchange-rate/internal/model"
	"testing"
)

func TestCreateSubscriber(t *testing.T) {
	testFilePath := "../../data/emails_test.txt"
	subscriberRepository := NewSubscriberFileRepository(testFilePath)

	testEmail := "test@gmail.com"
	testSubscriber := model.NewSubscriber(testEmail)

	if err := subscriberRepository.Create(testSubscriber); err != nil {
		t.Errorf("failed to add subscribers '%s' to file '%s': %v", testEmail, testFilePath, err)
	}

	testSubscribers, err := subscriberRepository.GetAll()
	if err != nil {
		t.Errorf("failed to get all subscribers '%s' from file '%s': %v", testEmail, testFilePath, err)
	}

	recordExists := false
	for _, subscriber := range testSubscribers {
		if subscriber.GetEmail() == testEmail {
			recordExists = true
		}
	}

	if recordExists == false {
		t.Errorf("subscriber '%s' doesn't exist in file '%s': %v", testEmail, testFilePath, err)
	}

	err = subscriberRepository.ClearFile()
	if err != nil {
		t.Errorf("failed to clear file '%s': %v", testFilePath, err)
	}
}
