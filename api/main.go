package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Membuat instance Fiber
	app := fiber.New()

	// Menyajikan file statis
	app.Static("/", "../build")

	// Menjalankan server di port 3000
	log.Println("Server berjalan di http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}