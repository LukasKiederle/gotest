package sender

import "github.com/LukasKiederle/gotest/src/exercises/3/email/mail"

type Sender interface {

	// Send an email to a given address with a  message.
	SendMail(address mail.Address, message string)
}
