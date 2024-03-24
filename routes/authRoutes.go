package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itzcodex24/edu-swipe-api/controllers"
)

func Auth(app *fiber.App) {
	app.Get("/", controllers.GetHello)
	app.Get("/api/user", controllers.User)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
}
