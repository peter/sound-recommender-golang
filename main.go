package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/peter/sound-recommender-golang/database"
	"github.com/peter/sound-recommender-golang/handlers"
	"github.com/peter/sound-recommender-golang/initializers"
)

func runInitializers() {
	log.Info("initialize starting")
	startTime := time.Now()

	database.Db = database.ConnectDB()

	initializers.MigrateModels(database.Db)

	initializers.DatabaseCrudTest(database.Db)

	elapsed := time.Since(startTime)
	log.Infof("initialize done - elapsed=%d", elapsed.Milliseconds())
}

func setupMiddleware(app *fiber.App) {
	app.Use(fiberLogger.New())
	// Recover from panic in any handler (https://docs.gofiber.io/guide/error-handling/)
	// TODO: could use a custom ErrorHandler
	app.Use(recover.New())
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Sound Recommender API")
	})

	// sound routes
	app.Post("/admin/sounds", handlers.CreateSound)
	app.Put("admin/sounds/:id", handlers.UpdateSound)
	app.Delete("admin/sounds/:id", handlers.DeleteSound)
	app.Get("sounds/:id", handlers.GetSound)
	app.Get("sounds", handlers.ListSounds)
}

func main() {
	runInitializers()

	app := fiber.New()

	setupMiddleware(app)

	setupRoutes(app)

	// https://apple.stackexchange.com/questions/393715/do-you-want-the-application-main-to-accept-incoming-network-connections-pop
	log.Fatal(app.Listen("localhost:8080"))
}
