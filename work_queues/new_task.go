package work_queues

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

// Send sends a message to the queue
func Send(body string) {
	// Create a connection to the server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue
	q, err := ch.QueueDeclare(
		"task_queue", // Name of the queue
		true,         // Durable
		false,        // Delete when unused
		false,        // Exclusive
		false,        // No-wait
		nil,          // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Publish a message
	err = ch.PublishWithContext(ctx,
		"",     // Exchange
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 2 = persistent
			ContentType:  "text/plain",    // MIME type
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
