package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/peter/sound-recommender-golang/database"
	"github.com/peter/sound-recommender-golang/helpers"
	"github.com/peter/sound-recommender-golang/models"
)

func CreateSound(c *fiber.Ctx) error {
	var sound models.Sound

	if err := c.BodyParser(&sound); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
	}
	if errs := helpers.InputValidator.Validate(sound); len(errs) > 0 && errs[0].Error {
		errors := make([]string, 0)
		for _, err := range errs {
			errors = append(errors, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(map[string]interface{}{"errors": errors})
	}

	// Gorm error handling: https://gorm.io/docs/error_handling.html
	if result := database.Db.Create(&sound); result.Error != nil {
		// TODO: should this error be handled and if so how?
		log.Error("CreateSound error", result.Error)
		return c.Status(500).JSON(map[string]interface{}{"error": "CreateSound error"})
	}

	return c.Status(200).JSON(sound)
}

func UpdateSound(c *fiber.Ctx) error {
	return c.SendString("TODO: UpdateSound")
	// id, err := c.ParamsInt("id")

	// var product models.Product

	// if err != nil {
	// 	return c.Status(400).JSON("Please ensure that :id is an integer")
	// }

	// err = findProduct(id, &product)

	// if err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }

	// type UpdateProduct struct {
	// 	Name         string `json:"name"`
	// 	SerialNumber string `json:"serial_number"`
	// }

	// var updateData UpdateProduct

	// if err := c.BodyParser(&updateData); err != nil {
	// 	return c.Status(500).JSON(err.Error())
	// }

	// product.Name = updateData.Name
	// product.SerialNumber = updateData.SerialNumber

	// database.Database.Db.Save(&product)

	// responseProduct := CreateResponseProduct(product)

	// return c.Status(200).JSON(responseProduct)
}

func DeleteSound(c *fiber.Ctx) error {
	return c.SendString("TODO: DeleteSound")
	// id, err := c.ParamsInt("id")

	// var user models.User

	// if err != nil {
	// 	return c.Status(400).JSON("Please ensure that :id is an integer")
	// }

	// err = findUser(id, &user)

	// if err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }

	// if err = database.Database.Db.Delete(&user).Error; err != nil {
	// 	return c.Status(404).JSON(err.Error())
	// }
	// return c.Status(200).JSON("Successfully deleted User")
}

func GetSound(c *fiber.Ctx) error {
	return c.SendString("TODO: GetSound")
	// id, err := c.ParamsInt("id")

	// var product models.Product

	// if err != nil {
	// 	return c.Status(400).JSON("Please ensure that :id is an integer")
	// }

	// if err := findProduct(id, &product); err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }

	// responseProduct := CreateResponseProduct(product)

	// return c.Status(200).JSON(responseProduct)
}

func ListSounds(c *fiber.Ctx) error {
	return c.SendString("TODO: ListSounds")
	// products := []models.Product{}
	// database.Database.Db.Find(&products)
	// responseProducts := []Product{}
	// for _, product := range products {
	// 	responseProduct := CreateResponseProduct(product)
	// 	responseProducts = append(responseProducts, responseProduct)
	// }

	// return c.Status(200).JSON(responseProducts)
}
