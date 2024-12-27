package services

import (
	"mime/multipart"
	"time"

	externalclients "github.com/uttom-akash/storage/internal/application/contracts/external_clients"
	"github.com/uttom-akash/storage/internal/application/contracts/services"
	file_domain "github.com/uttom-akash/storage/internal/domain/file"
)

type FileService struct {
	s3_client externalclients.IS3Client	
	rabbitmq_client externalclients.IRabbitMQClient
	file_repository file_domain.IFileRepository

}

func NewFileService(s3_client externalclients.IS3Client,
	 rabbitmq_client externalclients.IRabbitMQClient,
	 file_repository file_domain.IFileRepository) services.IFileService {

	fileService := &FileService{
		s3_client,
		rabbitmq_client,
		file_repository,
	}

	return fileService
}

func (documentHandler *FileService) UploadFile(file *multipart.FileHeader) {

	file_model := &file_domain.File{
		Name: file.Filename,
		CreatedAt: time.Now().String(),
	}

	documentHandler.file_repository.Create(file_model)

	go documentHandler.workInBackGround(file)
}


func (documentHandler *FileService) workInBackGround(file *multipart.FileHeader) {

	documentHandler.s3_client.UploadFile(file)

	documentHandler.rabbitmq_client.PublishMessage(file.Filename)
}
