package router

import (
	kelasRoutes "rest-api/internal/router/kelas"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")

	kelasRoutes.SetupKelasRoutes(api)
}
