package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uttom-akash/storage/internal/application/contracts/services"
)


type FileHandler struct {
	file_service services.IFileService
}

func NewFileHandler(server *gin.Engine, file_service services.IFileService) *FileHandler {
	fileHandler := &FileHandler{
		file_service,
	}

	documentHandlerRoute := server.Group("api/v1/storage/documents")

	documentHandlerRoute.POST("", fileHandler.upload)

	return fileHandler
}

func (fileHandler *FileHandler) upload(ctx *gin.Context) {

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

	fileHandler.file_service.UploadFile(file)

	fmt.Sprintf("file name: ", file.Filename)

	ctx.String(http.StatusOK, file.Filename)
}