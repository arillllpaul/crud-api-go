package handler

import (
    "encoding/json"
    "net/http"
    "log"
)

// Struktur Produk
type Produk struct {
    ID       int     `json:"id"`
    Nama     string  `json:"nama"`
    Harga    float64 `json:"harga"`
    Kategori string  `json:"kategori"`
}

// Handler untuk GetProduk
func GetProduk(w http.ResponseWriter, r *http.Request) {
    // Misalnya, ambil produk dari database
    produkList := []string{"Produk 1", "Produk 2", "Produk 3"} // Gantilah ini dengan logika untuk mengambil produk dari database

    // Mengirim response dalam format JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(produkList)
}

// Handler untuk CreateProduk
func CreateProduk(w http.ResponseWriter, r *http.Request) {
    var produk Produk
    err := json.NewDecoder(r.Body).Decode(&produk)
    if err != nil {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    // Simpan produk ke database (logika penyimpanan di sini)
    log.Println("Produk baru: ", produk)

    // Mengirim response sukses
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(produk)
}
