package main

import (
    "log"
    "net/http"
    "crud/handler"  // Pastikan path ini sesuai dengan nama modul dan struktur folder
    "github.com/gorilla/mux"
)



func main() {
    r := mux.NewRouter()

    // Rute untuk kategori
    r.HandleFunc("/kategori", handler.GetKategori).Methods("GET")
    r.HandleFunc("/kategori", handler.CreateKategori).Methods("POST")

    // Rute untuk produk
    r.HandleFunc("/produk", handler.GetProduk).Methods("GET")
    r.HandleFunc("/produk", handler.CreateProduk).Methods("POST")

    log.Println("Server berjalan di port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
