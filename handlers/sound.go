package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/peter/sound-recommender-golang/database"
	"github.com/peter/sound-recommender-golang/helpers"
	"github.com/peter/sound-recommender-golang/models"
)

func findSound(c *fiber.Ctx, id int, sound *models.Sound) bool {
	database.Db.Find(&sound, "id = ?", id)
	log.Info("pm debug findSound ", sound.ID, sound.ID == 0)
	if sound.ID == 0 {
		errorMessage := "Could not find sound"
		c.Status(fiber.StatusNotFound).JSON(helpers.ErrorResponseBody(errorMessage))
		return false
	}
	return true
}

func CreateSound(c *fiber.Ctx) error {
	var sound models.Sound

	if result := helpers.InputValidator.ValidateInput(c, &sound); result.Error != nil {
		return result.Error
	}

	if result := database.Db.Create(&sound); result.Error != nil {
		// Gorm error handling: https://gorm.io/docs/error_handling.html
		// TODO: should this error be handled here or in the error handler?
		// Can Gorm be configured to raise those errors (panic) automatically?
		// errorMessage := fmt.Sprintf("CreateSound database error: %s", result.Error)
		// log.Error(errorMessage)
		// return c.Status(500).JSON(helpers.ErrorResponseBody(errorMessage))
		panic(result.Error)
	}

	return c.Status(200).JSON(sound)
}

func UpdateSound(c *fiber.Ctx) error {
	return c.SendString("TODO: UpdateSound")
}

func DeleteSound(c *fiber.Ctx) error {
	return c.SendString("TODO: DeleteSound")
}

func GetSound(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(helpers.ErrorResponseBody("Please ensure that :id is an integer"))
	}

	var sound models.Sound

	if findResult := findSound(c, id, &sound); !findResult {
		return nil
	}

	return c.Status(200).JSON(sound)
}

func ListSounds(c *fiber.Ctx) error {
	sounds := []models.Sound{}
	database.Db.Find(&sounds)
	return c.Status(200).JSON(map[string]interface{}{"data": sounds})
}
