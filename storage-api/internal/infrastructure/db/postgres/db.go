package postgres

import (
	"log"

	file_domain "github.com/uttom-akash/storage/internal/domain/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	dsn := "user=username password=password dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to databae: %v", err)
	}

	return db
}

func createDatabase() {

}

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&file_domain.File{})

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}

func NewDB() (*gorm.DB, error) {
	db := connectDb()

	return db, nil
}
