package mail

import (
	"log"
)

type Address struct {
	Address string
}

type MailSenderImpl struct {
}

func (m *MailSenderImpl) SendMail(address Address, message string) {
	log.Println("Sending message with SMTP to " + address.Address + " message: " + message)
	return
}
