package main

import "fmt"

type SendInterface interface {
	SetReceiver(receiver string)
	SendMessage()
	Send()
}

type Sender struct {
	SendInterface
}

func NewSender(s SendInterface) *Sender {
	return &Sender{SendInterface: s}
}

type Email struct {
	receiver string
}

func (e *Email) SetReceiver(r string) {
	e.receiver = r
}

func (e *Email) SendMessage() {
	fmt.Printf("send email to %s\n", e.receiver)
}

func (e *Email) Send() {
	e.SetReceiver(e.receiver)
	e.SendMessage()
}

func main() {
	email := &Email{"abc@gmail.com"}
	s := NewSender(email)
	s.Send()
}
