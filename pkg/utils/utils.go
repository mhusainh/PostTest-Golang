package utils

import (
	"errors"
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
		RespondError(w, http.StatusInternalServerError, "Failed to encode JSON")
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

// ParseBody digunakan untuk memparsing body permintaan HTTP ke dalam struct yang diberikan
func ParseBody(r *http.Request, x interface{}) error {
    if r.Body == nil {
        return errors.New("request body is empty")
    }
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(x)
    if err != nil {
        return err
    }
    return nil
}