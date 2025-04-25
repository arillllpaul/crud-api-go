package main

import (
	"crud-uts/config"
	"crud-uts/handlers"
	"log"
	"net/http"
)

func main() {
	config.Connect()
	defer config.DB.Close()

	// Kategori routes
	http.HandleFunc("/kategori/create", handlers.CreateKategori)
	http.HandleFunc("/kategori/read", handlers.GetKategori)
	http.HandleFunc("/kategori/update/", handlers.UpdateKategori)
	http.HandleFunc("/kategori/delete/", handlers.DeleteKategori)

	// Produk routes
	http.HandleFunc("/produk/create", handlers.CreateProduk)
	http.HandleFunc("/produk/read", handlers.GetProduk)
	http.HandleFunc("/produk/update/", handlers.UpdateProduk)
	http.HandleFunc("/produk/delete/", handlers.DeleteProduk)

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}