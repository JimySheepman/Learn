package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	c, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Mail("sender@example.org"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("recipient@example.net"); err != nil {
		log.Fatal(err)
	}

	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
