// handlers/kategori_handler.go
package handlers

import (
	"crud-uts/config"
	"crud-uts/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func CreateKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var k models.Kategori
	if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if k.Nama == "" {
		http.Error(w, "Nama kategori tidak boleh kosong", http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec("INSERT INTO kategori(nama) VALUES(?)", k.Nama)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(k)
}

func GetKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT id, nama FROM kategori")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result []models.Kategori
	for rows.Next() {
		var k models.Kategori
		if err := rows.Scan(&k.ID, &k.Nama); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = append(result, k)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func DeleteKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/kategori/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var count int
	err = config.DB.QueryRow("SELECT COUNT(*) FROM produk WHERE kategori_id = ?", id).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Kategori tidak bisa dihapus karena masih digunakan oleh produk", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("DELETE FROM kategori WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Kategori deleted successfully",
		"id":      id,
	})
}
func UpdateKategori(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/kategori/update/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var k models.Kategori
	if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if k.Nama == "" {
		http.Error(w, "Nama kategori tidak boleh kosong", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("UPDATE kategori SET nama = ? WHERE id = ?", k.Nama, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(k)
}