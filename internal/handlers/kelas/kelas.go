package kelasHandler

import (
	"rest-api/database"
	"rest-api/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetKelass(c *fiber.Ctx) error {
	db := database.DB
	var kelas []model.Kelas

	// find all Kelas in the database
	db.Find(&kelas)

	// If no kelas is present return an error
	if len(kelas) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Kelas present", "data": nil})
	}

	// Else return Kelas
	return c.JSON(fiber.Map{"status": "success", "message": "Kelas Found", "data": kelas})
}

func CreateKelas(c *fiber.Ctx) error {
	db := database.DB
	kelas := new(model.Kelas)

	// Store the body in the kelas and return error if encountered
	err := c.BodyParser(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the kelas
	kelas.ID = uuid.New()
	// Create the kelas and return error if encountered
	err = db.Create(&kelas).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create kelas", "data": err})
	}

	// Return the created kelas
	return c.JSON(fiber.Map{"status": "success", "message": "Created kelas", "data": kelas})
}

func GetKelas(c *fiber.Ctx) error {
	db := database.DB
	var kelas model.Kelas

	// Read the param KelasId
	id := c.Params("KelasId")

	// Find the Kelas with the given Id
	db.Find(&kelas, "id = ?", id)

	// If no such Kelas present return an error
	if kelas.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Kelas present", "data": nil})
	}

	// Return the Kelas with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Kelass Found", "data": kelas})
}

func UpdateKelas(c *fiber.Ctx) error {
	type updateKelas struct {
		Name      string `json:"name"`
		WaliKelas string `json:"wali_kelas"`
	}
	db := database.DB
	var kelas model.Kelas

	// Read the param ID
	id := c.Params("ID")

	// Find the Kelas with the given Id
	db.Find(&kelas, "id = ?", id)

	// If no such Kelas present return an error
	if kelas.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Kelas present", "data": nil})
	}

	// Store the body containing the updated data and return error if encountered
	var updateKelasData updateKelas
	err := c.BodyParser(&updateKelasData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the Kelas
	kelas.Name = updateKelasData.Name
	kelas.WaliKelas = updateKelasData.WaliKelas

	// Save the Changes
	db.Save(&kelas)

	// Return the updated Kelas
	return c.JSON(fiber.Map{"status": "success", "message": "Kelass Found", "data": kelas})
}

func DeleteKelas(c *fiber.Ctx) error {
	db := database.DB
	var kelas model.Kelas

	// Read the param KelasId
	id := c.Params("ID")

	// Find the Kelas with the given Id
	db.Find(&kelas, "id = ?", id)

	// If no such Kelas present return an error
	if kelas.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Kelas present", "data": nil})
	}

	// Delete the Kelas and return error if encountered
	err := db.Delete(&kelas, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete Kelas", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Kelas"})
}
