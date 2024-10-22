package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandleErr digunakan untuk mencetak error dan menghentikan aplikasi jika ada error kritis
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err) // Tampilkan pesan error dan hentikan aplikasi
	}
}

// RespondJSON digunakan untuk mengirim respons dalam format JSON
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload) // Ubah payload menjadi JSON
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") // Atur header respons
	w.WriteHeader(status)                              // Kirim status kode HTTP
	w.Write(response)                                  // Kirim data JSON sebagai respons
}

// RespondError digunakan untuk mengirim respons error dalam format JSON
func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, map[string]string{"error": message}) // Kirim pesan error dalam JSON
}
