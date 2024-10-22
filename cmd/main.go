package main

import (
	"log"
	"net/http"
	"Weekly-Task3/pkg/config"
	"Weekly-Task3/pkg/routes"
)

func main() {
	db := config.ConnectDB() // Hubungkan ke database
	defer func() {
		if err := db.Close(); err != nil { // Tutup koneksi dan cek error
			log.Fatalf("Error closing database: %v", err)
		}
	}()

	router := routes.SetupRoutes() // Siapkan semua rute

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil { // Jalankan server dan cek error
		log.Fatalf("Error starting server: %v", err)
	}
}
