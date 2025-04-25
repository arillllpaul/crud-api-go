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
    // Contoh data kategori yang lebih banyak
    kategoriList := []Kategori{
        {ID: 1, Nama: "Elektronik"},
        {ID: 2, Nama: "Pakaian"},
        {ID: 3, Nama: "Buku"},
        {ID: 4, Nama: "Alat Tulis"},
        {ID: 5, Nama: "Peralatan Dapur"},
        {ID: 6, Nama: "Makanan & Minuman"},
        {ID: 7, Nama: "Kesehatan"},
        {ID: 8, Nama: "Olahraga"},
        {ID: 9, Nama: "Kecantikan"},
        {ID: 10, Nama: "Mainan Anak"},
    }

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
    log.Println("Kategori baru ditambahkan:", kategori)

    // Mengirim response sukses
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(kategori)
}
