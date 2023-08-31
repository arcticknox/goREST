package main

import (
	"goREST/db"
	"goREST/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// db
	db.InitDB(utils.GetConfig("DB_NAME"))
	app.Listen(":3000")
}
