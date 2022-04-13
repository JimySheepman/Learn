package main

import (
	"log"
	"net/smtp"
)

var (
	from       = "gopher@example.net"
	msg        = []byte("dummy message")
	recipients = []string{"foo@example.com"}
)

func main() {
	hostname := "mail.example.com"
	auth := smtp.PlainAuth("", "user@example.com", "password", hostname)

	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}
