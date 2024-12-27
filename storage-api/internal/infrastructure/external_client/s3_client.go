package externalclient

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	externalclients "github.com/uttom-akash/storage/internal/application/contracts/external_clients"
)

type S3Client struct {
	config *ExternalClientConfig
}

func NewS3Client(config *ExternalClientConfig) externalclients.IS3Client {
	return &S3Client{config}
}


func (s *S3Client) UploadFile(file *multipart.FileHeader) {
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
}