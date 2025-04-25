// handlers/produk_handler.go
package handlers

import (
	"crud-uts/config"
	"crud-uts/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CreateProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p models.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if p.Nama == "" {
		http.Error(w, "Nama produk tidak boleh kosong", http.StatusBadRequest)
		return
	}
	if p.Harga <= 0 {
		http.Error(w, "Harga harus lebih dari 0", http.StatusBadRequest)
		return
	}

	// Cek keberadaan kategori
	var kategoriExists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM kategori WHERE id = ?)", p.KategoriID).Scan(&kategoriExists)
	if err != nil {
		log.Printf("Error checking kategori existence: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !kategoriExists {
		http.Error(w, "Kategori tidak ditemukan", http.StatusBadRequest)
		return
	}

	// Insert produk
	res, err := config.DB.Exec("INSERT INTO produk(nama, harga, kategori_id) VALUES(?, ?, ?)",
		p.Nama, p.Harga, p.KategoriID)
	if err != nil {
		log.Printf("Error inserting produk: %v", err)
		http.Error(w, "Gagal membuat produk", http.StatusInternalServerError)
		return
	}

	// Ambil ID yang baru dibuat
	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
	} else {
		p.ID = int(id)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func GetProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query(`
		SELECT p.id, p.nama, p.harga, p.kategori_id, k.nama as kategori_nama 
		FROM produk p
		JOIN kategori k ON p.kategori_id = k.id
		ORDER BY p.id
	`)
	if err != nil {
		log.Printf("Error querying produk: %v", err)
		http.Error(w, "Gagal mengambil data produk", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result []models.ProdukWithKategori
	for rows.Next() {
		var p models.ProdukWithKategori
		if err := rows.Scan(&p.ID, &p.Nama, &p.Harga, &p.KategoriID, &p.KategoriNama); err != nil {
			log.Printf("Error scanning produk row: %v", err)
			http.Error(w, "Gagal memproses data produk", http.StatusInternalServerError)
			return
		}
		result = append(result, p)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error after iterating rows: %v", err)
		http.Error(w, "Gagal memproses data produk", http.StatusInternalServerError)
		return
	}

	if len(result) == 0 {
		result = []models.ProdukWithKategori{} // Return empty array instead of null
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/produk/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID produk tidak valid", http.StatusBadRequest)
		return
	}

	// Cek apakah produk ada sebelum menghapus
	var exists bool
	err = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM produk WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		log.Printf("Error checking produk existence: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
		return
	}

	_, err = config.DB.Exec("DELETE FROM produk WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting produk: %v", err)
		http.Error(w, "Gagal menghapus produk", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Produk berhasil dihapus",
		"id":      id,
	})
}

func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/produk/update/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID produk tidak valid", http.StatusBadRequest)
		return
	}

	var p models.Produk
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if p.Nama == "" {
		http.Error(w, "Nama produk tidak boleh kosong", http.StatusBadRequest)
		return
	}
	if p.Harga <= 0 {
		http.Error(w, "Harga harus lebih dari 0", http.StatusBadRequest)
		return
	}

	// Cek apakah produk ada
	var produkExists bool
	err = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM produk WHERE id = ?)", id).Scan(&produkExists)
	if err != nil {
		log.Printf("Error checking produk existence: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !produkExists {
		http.Error(w, "Produk tidak ditemukan", http.StatusNotFound)
		return
	}

	// Cek keberadaan kategori
	var kategoriExists bool
	err = config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM kategori WHERE id = ?)", p.KategoriID).Scan(&kategoriExists)
	if err != nil {
		log.Printf("Error checking kategori existence: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if !kategoriExists {
		http.Error(w, "Kategori tidak ditemukan", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("UPDATE produk SET nama = ?, harga = ?, kategori_id = ? WHERE id = ?",
		p.Nama, p.Harga, p.KategoriID, id)
	if err != nil {
		log.Printf("Error updating produk: %v", err)
		http.Error(w, "Gagal memperbarui produk", http.StatusInternalServerError)
		return
	}

	// Set ID untuk response
	p.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}