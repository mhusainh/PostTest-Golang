package config

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Driver SQLite
	"Weekly-Task3/pkg/models"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "bookstore.db") // Buat koneksi ke SQLite
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Tangkap error jika koneksi gagal
	}

	if err := db.AutoMigrate(&models.Book{}).Error; err != nil { // Migrasi dan cek error
		log.Fatalf("Migration failed: %v", err)
	}

	return db
}
