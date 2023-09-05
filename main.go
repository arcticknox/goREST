package main

import (
	"goREST/db"
	"goREST/routes"
	"goREST/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize db
	db.InitDB(utils.GetConfig("DB_NAME"))
	// Load routes
	routes.LoadRoutes(app)
	app.Listen(":3000")
}
