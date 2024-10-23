package config

import (
	"log"

	"Weekly-Task3/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{}) // Buat koneksi ke SQLite
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Tangkap error jika koneksi gagal
	}

	if err := db.AutoMigrate(&models.Book{}); err != nil { // Migrasi dan cek error
		log.Fatalf("Migration failed: %v", err)
	}

	return db
}
