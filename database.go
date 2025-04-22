package main

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
    dsn := "root:password@tcp(localhost:3306)/nama_database"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error connecting to database: ", err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal("Database connection failed: ", err)
    }
    log.Println("Koneksi ke database berhasil")
    return db
}
