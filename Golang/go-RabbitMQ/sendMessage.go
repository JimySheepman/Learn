// Publisher
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// Set the way error message are displayed in the terminal
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ or send a message,
	// if there are any errors connecting
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Create a queue to send the message to.
	q, err := ch.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // dele when unsused
		false,          //exclusive
		false,          // no-wait
		nil,            // arguments

	)
	failOnError(err, "Failed to declare a queue")

	for i := 3; i < 103; i++ {

		// set the payload for the message.
		body := fmt.Sprintf("Golang is awesome - Keep Moving Forward!-%d", i)
		time.Sleep(250 * time.Millisecond)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		// If there is an error publishing the message,
		// a log will be displayed in the terminal.
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Congrats, sending message: %s", body)
	}
}
