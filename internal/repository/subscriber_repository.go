package repository

import (
	"bitcoin-exchange-rate/internal/model"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrEmailAlreadyExist = errors.New("email already exists in the file")

type SubscriberFileRepository struct {
	filePath string
}

func NewSubscriberFileRepository(filePath string) *SubscriberFileRepository {
	return &SubscriberFileRepository{
		filePath: filePath,
	}
}

func (r *SubscriberFileRepository) GetAll() ([]*model.Subscriber, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var subscribers []*model.Subscriber

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		subscribers = append(subscribers, model.NewSubscriber(strings.TrimSpace(scanner.Text())))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}

func (r *SubscriberFileRepository) Create(subscriber *model.Subscriber) error {
	content, err := os.ReadFile(r.filePath)
	if err != nil {
		return err
	}

	if strings.Contains(string(content), subscriber.GetEmail()) {
		fmt.Println("test: ", subscriber)
		fmt.Println(string(content))
		return ErrEmailAlreadyExist
	}

	file, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString("\n" + subscriber.GetEmail()); err != nil {
		return err
	}

	fmt.Println("Subscriber appended to the file.")
	return nil
}

func (r *SubscriberFileRepository) ClearFile() error {
	file, err := os.OpenFile(r.filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
