package model

type Subscriber struct {
	email string
}

func NewSubscriber(email string) *Subscriber {
	return &Subscriber{
		email: email,
	}
}

func (s Subscriber) GetEmail() string {
	return s.email
}
