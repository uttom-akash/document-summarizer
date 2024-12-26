package postgres

import (
	"log"

	"github.com/uttom-akash/storage/internal/domain/file"
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

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&file.File{})

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}

func NewDB() (*gorm.DB, error) {
	db := connectDb()

	runMigrations(db)

	return db, nil
}
