package main

import (
	"log"
	"os"
	"user-management/config"
	"user-management/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	config.SetupLog()

	// Inisialisasi Gin
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		panic("Error loading .env file")
	}

	// Initialize database connection
	config.Connect()

	// Setup routes
	routes.SetupRoutes(r)

	// Menentukan port service
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8080" // Default port jika tidak ada di env
	}

	// Jalankan server
	r.Run(":" + port)
	log.Println("Server berjalan di port : ", port)
}
