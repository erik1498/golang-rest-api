package main

import (
	"rest-api/database"
	"rest-api/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	router.SetupRoutes(app)

	app.Listen(":3000")
}
