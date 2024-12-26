package services

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type DocumentHandler struct {
	// uploaderChan chan *multipart.FileHeader
}

func NewDocumentHandler(server *gin.Engine) *DocumentHandler {
	documentHandler := &DocumentHandler{
		// uploaderChan: make(chan *multipart.FileHeader, 100)
	}

	documentHandlerRoute := server.Group("api/v1/storage/documents")

	documentHandlerRoute.POST("", documentHandler.upload)

	return documentHandler
}

func (documentHandler *DocumentHandler) upload(ctx *gin.Context) {

	if ctx.ContentType() != "multipart/form-data" {
		log.Println("Invalid Content-Type. Expected multipart/form-data.")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be multipart/form-data"})
		return
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		log.Fatal("file error", err)
		ctx.String(http.StatusInternalServerError, "")
		return
	}

	go documentHandler.uploadWorker(file)

	fmt.Sprintf("file name: ", file.Filename)

	ctx.String(http.StatusOK, file.Filename)
}

func (documentHandler *DocumentHandler) uploadWorker(file *multipart.FileHeader) {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	cfg.Region = "eu-north-1"

	s3Client := s3.NewFromConfig(cfg)

	body, err := file.Open()

	if err != nil {
		log.Fatal(err)
	}

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("web-bucket-8a9e1fc"),
		Key:    aws.String(file.Filename),
		Body:   body,
	})

	if err != nil {
		log.Fatalf("failed to upload file, %v", err)
	}

	sendMessage(file.Filename)

	log.Printf("File uploaded successfully to S3 bucket: , with key: \n")
}

func sendMessage(filename string) {
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
