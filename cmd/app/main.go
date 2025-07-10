package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/lucaslimafernandes/rinha-backend-2025/pkg/router"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file", err)
	}

}

func main() {

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))

}
