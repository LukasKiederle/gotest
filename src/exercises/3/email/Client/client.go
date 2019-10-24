package main

import (
	"github.com/LukasKiederle/gotest/src/exercises/3/email/mail"
	sender "github.com/LukasKiederle/gotest/src/exercises/3/email/senderInterface"
	"github.com/LukasKiederle/gotest/src/exercises/3/util"
)

var Registry = util.NewRegistry()

func sendTestMail(message string) {
	// Create an implementation for the mail.Sender interface
	sender := Registry.Get("sender.Sender").(sender.Sender)

	mailaddrs := mail.Address{Address: "test@test.com"}
	sender.SendMail(mailaddrs, message)
}

func main() {
	sendTestMail("test")
}
