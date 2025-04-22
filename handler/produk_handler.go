package handler

import (
    "crud/model" // Impor model Produk dari package model
    "net/http"
    "encoding/json"
    "log"
)

// Handler untuk GetProduk - Menampilkan daftar produk
func GetProduk(w http.ResponseWriter, r *http.Request) {
    // Misalnya, ambil produk dari database
    produkList := []model.Produk{
        {ID: 1, Nama: "Produk 1", Harga: 100.0, Kategori: "Kategori A"},
        {ID: 2, Nama: "Produk 2", Harga: 200.0, Kategori: "Kategori B"},
    }

    // Mengirim response dalam format JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(produkList)
}

// Handler untuk CreateProduk - Menambahkan produk baru
func CreateProduk(w http.ResponseWriter, r *http.Request) {
    var produk model.Produk
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
