package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Variabel global untuk instance Database
var DB *gorm.DB

// ConnectDatabase berfungsi untuk menginisialisasi koneksi SQLite (Atau MySQL/Postgres nantinya)
func ConnectDatabase(dbName string) {
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database!", err)
	}

	DB = database
	log.Println("Database berhasil terkoneksi")
}
