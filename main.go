package main

import (
	"api-storage-github/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No possible load environment variables")
	}

	app := fiber.New()

	app.Get("/execute-codes/:scriptName", handlers.ExecuteCodefunc)
	app.Post("/codes", handlers.SaveCode)
	app.Listen(":3000")
}
