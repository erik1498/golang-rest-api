package kelasRoutes

import (
	kelasHandler "rest-api/internal/handlers/kelas"

	"github.com/gofiber/fiber/v2"
)

func SetupKelasRoutes(router fiber.Router) {
	kelas := router.Group("/kelas")

	kelas.Get("/", kelasHandler.GetKelass)
	kelas.Get("/:ID", kelasHandler.GetKelas)
	kelas.Post("/", kelasHandler.CreateKelas)
	kelas.Put(":/ID", kelasHandler.UpdateKelas)
	kelas.Delete("/:ID", kelasHandler.DeleteKelas)
}
