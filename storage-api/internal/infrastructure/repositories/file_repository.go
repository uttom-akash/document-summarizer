package repositories

import (
	"github.com/uttom-akash/storage/internal/domain/file"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db}
}

func (repo *FileRepository) Create(file *file.File) error {
	result := repo.db.Create(file)
	return result.Error
}
