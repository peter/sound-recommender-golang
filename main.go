package main

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/peter/sound-recommender-golang/initializers"
	"github.com/peter/sound-recommender-golang/models"
)

func main() {
	db := initializers.ConnectDB()
	db.AutoMigrate(&models.Sound{})
	initializers.DatabaseCrudTest(db)

	app := fiber.New()

	// Middleware
	app.Use(fiberLogger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// https://apple.stackexchange.com/questions/393715/do-you-want-the-application-main-to-accept-incoming-network-connections-pop
	app.Listen("localhost:8080")
}
