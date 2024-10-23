package controllers

import (
	"net/http"

	"Weekly-Task3/pkg/models"
	"Weekly-Task3/pkg/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB // Variabel global untuk koneksi database

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	if err := DB.Find(&books).Error; err != nil { // Cek apakah ada error saat mengambil data
		utils.RespondError(w, http.StatusInternalServerError, "Unable to retrieve books")
		return
	}
	utils.RespondJSON(w, http.StatusOK, books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := DB.First(&book, params["bookid"]).Error; err != nil { // Cek error saat mengambil buku
		utils.RespondError(w, http.StatusNotFound, "Book not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := utils.ParseBody(r, &book); err != nil { // Menggunakan utils untuk parsing body
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}
	if err := DB.Create(&book).Error; err != nil { // Cek error saat menyimpan buku
		utils.RespondError(w, http.StatusInternalServerError, "Unable to create book")
		return
	}
	utils.RespondJSON(w, http.StatusCreated, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var existingBook models.Book
	if err := DB.First(&existingBook, params["bookid"]).Error; err != nil { // Cek error jika buku tidak ditemukan
		utils.RespondError(w, http.StatusNotFound, "Book not found")
		return
	}

	var updatedBook models.Book
	if err := utils.ParseBody(r, &updatedBook); err != nil { // Menggunakan utils untuk parsing body
		utils.RespondError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Update field-field buku
	existingBook.Title = updatedBook.Title
	existingBook.Author = updatedBook.Author
	existingBook.Price = updatedBook.Price

	if err := DB.Save(&existingBook).Error; err != nil { // Cek error saat menyimpan perubahan
		utils.RespondError(w, http.StatusInternalServerError, "Unable to update book")
		return
	}
	utils.RespondJSON(w, http.StatusOK, existingBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := DB.First(&book, params["bookid"]).Error; err != nil { // Cek jika buku tidak ditemukan
		utils.RespondError(w, http.StatusNotFound, "Book not found")
		return
	}
	if err := DB.Delete(&book).Error; err != nil { // Cek error saat menghapus buku
		utils.RespondError(w, http.StatusInternalServerError, "Unable to delete book")
		return
	}
	w.WriteHeader(http.StatusNoContent) // Kirim status 204 No Content jika berhasil
}
