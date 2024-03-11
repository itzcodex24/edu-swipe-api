package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itzcodex24/edu-swipe-api/database"
	"github.com/itzcodex24/edu-swipe-api/routes"
)

func main() {

	database.Connect()
	app := fiber.New()

	routes.Auth(app)
	if err := app.Listen(":3001"); err != nil {
		if err = fmt.Errorf("error: %v", err); err != nil {
			panic(err)
		}
	}
}
