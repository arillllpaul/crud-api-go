package handler

import (
    "crud/model" // Impor model Kategori dari package model
    "net/http"
    "encoding/json"
    "log"
)

// Handler untuk GetKategori - Menampilkan daftar kategori
func GetKategori(w http.ResponseWriter, r *http.Request) {
    // Misalnya, ambil kategori dari database
    kategoriList := []model.Kategori{
        {ID: 1, Nama: "Kategori A"},
        {ID: 2, Nama: "Kategori b"},
    }

    // Mengirim response dalam format JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(kategoriList)
}

// Handler untuk CreateKategori - Menambahkan kategori baru
func CreateKategori(w http.ResponseWriter, r *http.Request) {
    var kategori model.Kategori
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

// Handler untuk UpdateKategori - Mengubah data kategori
func UpdateKategori(w http.ResponseWriter, r *http.Request) {
    var kategori model.Kategori
    err := json.NewDecoder(r.Body).Decode(&kategori)
    if err != nil {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    // Update kategori di database (logika update di sini)
    log.Println("Kategori diupdate: ", kategori)

    // Mengirim response sukses
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(kategori)
}

// Handler untuk DeleteKategori - Menghapus kategori
func DeleteKategori(w http.ResponseWriter, r *http.Request) {
    var kategori model.Kategori
    err := json.NewDecoder(r.Body).Decode(&kategori)
    if err != nil {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    // Hapus kategori dari database (logika penghapusan di sini)
    log.Println("Kategori dihapus: ", kategori)

    // Mengirim response sukses
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(kategori)
}
