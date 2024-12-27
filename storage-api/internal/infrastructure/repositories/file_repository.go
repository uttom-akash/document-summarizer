package repositories

import (
	file_domain "github.com/uttom-akash/storage/internal/domain/file"
	"gorm.io/gorm"
)

type FileRepository struct {
	repository *genericRepository[file_domain.File]
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{repository: NewGenericRepository[file_domain.File](db)}
}

func (repo *FileRepository) Create(file *file_domain.File) error {
	result := repo.repository.db.Create(file)
	return result.Error
}
