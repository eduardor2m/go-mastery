package main

import "log"

type Message interface {
	Send(to string, message string) error
}

type Email struct{}

func (e *Email) Send(to string, message string) error {
	log.Println("Sending email to", to, "with message:", message)
	return nil
}

type SMS struct{}

func (s *SMS) Send(to string, message string) error {
	log.Println("Sending SMS to", to, "with message:", message)
	return nil
}

type Notification struct {
	message Message
}

func (n *Notification) Send(to string, message string) error {
	return n.message.Send(to, message)
}

func main() {
	email := &Email{}
	sms := &SMS{}

	notification := &Notification{email}
	notification.Send("deveduardomelo@gmail.com", "Hello, World!")

	notification = &Notification{sms}
	notification.Send("123456789", "Hello, World!")
}
