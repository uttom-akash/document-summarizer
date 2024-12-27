package externalclients

import "mime/multipart"

type IS3Client interface {
	UploadFile(file *multipart.FileHeader)
}