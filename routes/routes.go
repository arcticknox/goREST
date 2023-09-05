package routes

import (
	"goREST/handlers"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	app.Post("/api/v1/signup", handlers.SignUp)
	app.Post("/api/v1/login", handlers.Login)
}
