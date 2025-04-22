package handler

import (
    "encoding/json"
    "net/http"
    "log"
)

// Struktur Kategori
type Kategori struct {
    ID   int    `json:"id"`
    Nama string `json:"nama"`
}

// Handler untuk GetKategori
func GetKategori(w http.ResponseWriter, r *http.Request) {
    // Misalnya, ambil kategori dari database
    kategoriList := []string{"Kategori 1", "Kategori 2", "Kategori 3"} // Gantilah ini dengan logika untuk mengambil kategori dari database

    // Mengirim response dalam format JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(kategoriList)
}

// Handler untuk CreateKategori
func CreateKategori(w http.ResponseWriter, r *http.Request) {
    var kategori Kategori
    err := json.NewDecoder(r.Body).Decode(&kategori)
    if err != nil {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    // Simpan kategori ke database (logika penyimpanan di sini)
    log.Println("Kategori baru: ", kategori)

    // Mengirim response sukses
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(kategori)
}
