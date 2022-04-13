package main

import (
	"log"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
