package file

type IFileRepository interface {
	Create(file *File) error
	GetById(id string) (*File, error)
	Delete(id string) error
}
