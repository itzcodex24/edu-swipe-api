package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itzcodex24/edu-swipe-api/database"
)

func main() {

	database.Connect()
	app := fiber.New()

	if err := app.Listen(":3001"); err != nil {
		fmt.Errorf("error: %v", err)
	}
}
