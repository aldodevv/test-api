package main

import (
	"log"
	"os"

	"advanced/config"
	"advanced/controllers"
	"advanced/middlewares"
	"advanced/structs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load file variabel '.env'
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan variabel OS default")
	}

	// 2. Setup Koneksi Database (Konfigurasi URL ada di dalam ConnectDatabase)
	config.ConnectDatabase()

	// AutoMigrate: GORM akan otomatis membuat tabel di database 
	// yang mensinkronisasi field-field pada Model User secara otomatis.
	log.Println("Menjalankan Auto Migrasi Database...")
	config.DB.AutoMigrate(&structs.User{})

	// 3. Setup Framework Gin (Router)
	r := gin.Default()

	// 4. Setup Routes Config
	// Menggunakan grouping route agak rapih dan bisa dikasih middleware di level group
	api := r.Group("/api")
	{
		// Endpoint publik (Tidak butuh token)
		api.POST("/users/register", controllers.CreateUser)

		// Endpoint terproteksi (Butuh token)
		// Kita sisipkan middlware "SimpleAuthMiddleware" sebelum mengeksekusi getAllUsers
		secured := api.Group("/secured")
		secured.Use(middlewares.SimpleAuthMiddleware())
		{
			secured.GET("/users", controllers.GetAllUsers)
		}
	}

	// 5. Jalankan Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server berjalan di port http://localhost:%s", port)
	r.Run(":" + port)
}
