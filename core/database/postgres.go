package database

import (
	"log"

	"advanced/core/environment"
	"advanced/structs"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProvidePostgreSQL adalah Injector Component yang memerlukan modul (Config)
// Tetap dipertahankan namanya agar tidak perlu re-generate Google Wire
func ProvidePostgreSQL(cfg *environment.Config) (*gorm.DB, error) {
	dsn := cfg.DATABASE_URL
	var db *gorm.DB
	var err error

	// Mengakomodasi fallback ke SQLite jika koneksi PostgreSQL belum di-setup di lokal
	if dsn == "app.db" || dsn == "" {
		log.Println("Menggunakan SQLite lokal karena ekstensi DATABASE_URL = app.db ...")
		db, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	} else {
		log.Println("Menggunakan PostgreSQL Remote/Lokal ...")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	// Auto migrate
	_ = db.AutoMigrate(&structs.User{})

	return db, nil
}
