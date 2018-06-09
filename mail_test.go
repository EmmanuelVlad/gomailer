package gomailer

import (
	"testing"
)

func TestMail(t *testing.T) {
	mail := New()
	mail.Remote = "localhost:25"
	mail.Sender = "sender@gomailer.com"
	mail.Recipient = "recipient@gomailer.com"
	mail.Content = "example"
	if err := mail.Send(); err != nil {
		t.Error(err)
	}
}
