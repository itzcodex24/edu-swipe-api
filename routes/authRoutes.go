package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itzcodex24/edu-swipe-api/controllers"
)

func Auth(app *fiber.App) {
	app.Get("/", controllers.GetHello)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Post("/api/register", controllers.Register)
	app.Get("/api/user", controllers.User)
}
