package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itzcodex24/edu-swipe-api/database"
)

func main() {
	app := fiber.New()

	if err := app.Listen(":3001"); err != nil {
		panic("Failed to start the server..")
	}

	database.Connect()
}
