package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	// Konfigurasi untuk MAMP
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/uts_web?parseTime=true")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	// Tambahkan ping untuk memastikan koneksi berhasil
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}

	log.Println("Successfully connected to database")
}