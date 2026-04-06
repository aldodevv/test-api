package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variabel global untuk instance Database
var DB *gorm.DB

// ConnectDatabase berfungsi untuk menginisialisasi koneksi PostgreSQL secara aman lewat Environment
func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Gagal konek DB: Variabel DATABASE_URL tidak ditemukan di environment (.env/fly.io)!")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database!", err)
	}

	DB = database
	log.Println("Database berhasil terkoneksi")
}
