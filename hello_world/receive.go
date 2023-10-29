package hello_world

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Receive() {
	// Create a connection to the server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	// Declare a queue
	q, err := ch.QueueDeclare(
		"hello", // Name of the queue
		false,   // Durable
		false,   // Delete when unused
		false,   // Exclusive
		false,   // No-wait
		nil,     // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Consume a message
	msgs, err := ch.Consume(
		q.Name, // Name of the queue
		"",     // Consumer
		true,   // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	failOnError(err, "Failed to register a consumer")

	// Create a channel to receive messages
	forever := make(chan bool)

	// Create a goroutine
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
