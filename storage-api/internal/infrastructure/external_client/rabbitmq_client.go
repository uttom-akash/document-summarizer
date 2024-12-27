package externalclient

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
	externalclients "github.com/uttom-akash/storage/internal/application/contracts/external_clients"
)

type RabbitMQClient struct {
	config *ExternalClientConfig
}


func NewRabbitMQClient(config *ExternalClientConfig) externalclients.IRabbitMQClient {
	return &RabbitMQClient{ config}
}

func (client *RabbitMQClient)PublishMessage(filename string) {
	// Connect to RabbitMQ server
	rabbit_mq_url := os.Getenv("RABBIT_MQ_URL")
	conn, err := amqp.Dial(rabbit_mq_url) // Adjust RabbitMQ URI if necessary
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a new channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	queue, err := ch.QueueDeclare(
		"my_queue", // Name of the queue
		true,       // Durable, so that the queue survives RabbitMQ restarts
		false,      // Delete when unused
		false,      // Exclusive to the connection
		false,      // No-wait
		nil,        // Arguments
	)
	if err != nil {
		log.Fatalf("failed to declare a queue: %v", err)
	}

	// Send a message to the queue
	body := filename
	err = ch.Publish(
		"",         // Exchange
		queue.Name, // Routing key (queue name)
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("failed to publish a message: %v", err)
	}

	fmt.Println("Message sent to queue:", queue.Name)
}
