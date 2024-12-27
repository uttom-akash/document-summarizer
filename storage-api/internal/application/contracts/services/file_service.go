package services

import "mime/multipart"

type IFileService interface {
	UploadFile(file *multipart.FileHeader)
}