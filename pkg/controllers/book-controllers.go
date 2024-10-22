package controllers

import (
	"encoding/json"
	"net/http"
	"Weekly-Task3/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB // Variabel global untuk koneksi database

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	if err := DB.Find(&books).Error; err != nil { // Cek apakah ada error saat mengambil data
		http.Error(w, "Unable to retrieve books", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(books); err != nil { // Cek apakah ada error saat mengirim JSON
		http.Error(w, "Unable to encode books to JSON", http.StatusInternalServerError)
	}
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := DB.First(&book, params["bookid"]).Error; err != nil { // Cek error saat mengambil buku
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(book); err != nil { // Cek error saat mengirim JSON
		http.Error(w, "Unable to encode book to JSON", http.StatusInternalServerError)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := DB.First(&book, params["bookid"]).Error; err != nil { // Cek error jika buku tidak ditemukan
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil { // Cek error saat mendekode JSON
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := DB.Save(&book).Error; err != nil { // Cek error saat menyimpan perubahan
		http.Error(w, "Unable to update book", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(book); err != nil { // Cek error saat mengirim JSON
		http.Error(w, "Unable to encode book to JSON", http.StatusInternalServerError)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := DB.First(&book, params["bookid"]).Error; err != nil { // Cek jika buku tidak ditemukan
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if err := DB.Delete(&book).Error; err != nil { // Cek error saat menghapus buku
		http.Error(w, "Unable to delete book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // Kirim status 204 No Content jika berhasil
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil { // Cek error saat mendekode JSON
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := DB.Create(&book).Error; err != nil { // Cek error saat menyimpan buku
		http.Error(w, "Unable to create book", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(book); err != nil { // Cek error saat mengirim JSON
		http.Error(w, "Unable to encode book to JSON", http.StatusInternalServerError)
	}
}
