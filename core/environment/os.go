package environment

import (
	"os"

	"github.com/joho/godotenv"
)

// Config Struktur penampung seluruh isian .env
type Config struct {
	PORT         string
	DATABASE_URL string
	SECRET_KEY   string
}

// ProvideConfig adalah Injector Component untuk google/wire
func ProvideConfig() (*Config, error) {
	// Menjalankan auto-load layaknya Odin
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		PORT:         port,
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		SECRET_KEY:   os.Getenv("SECRET_KEY"),
	}, nil
}
