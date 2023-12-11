package main

import (
	"fmt"

	"encoding/json"

	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/peter/sound-recommender-golang/models"
)

func printJson(prefix string, obj any) {
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prefix, string(jsonBytes))
}

func printStruct(prefix string, obj any) {
	fmt.Printf("%s %+v\n", prefix, obj)
}

func dbConnect() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=sound-recommender port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Verbose logging of SQL queries with Info level
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func dbCrudTest(db *gorm.DB) {
	// CREATE
	sound := models.Sound{
		Title:  "Foobar",
		Genres: []string{"dance pop", "latin"},
		Meta: datatypes.JSONMap{
			"credits": []map[string]interface{}{
				{
					"name": "Elton John",
					"role": "ARTIST",
				},
			},
		},
	}
	result := db.Create(&sound)
	if result.Error != nil {
		fmt.Println("Create error", result.Error)
	}
	fmt.Println("sound.ID", sound.ID)
	printJson("created sound", sound)

	// FIND BY ID
	var findSound models.Sound
	db.First(&findSound, sound.ID)
	printJson("find by id sound", findSound)

	// UPDATE
	// https://gorm.io/docs/update.html
	// db.Model(&sound).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	db.Model(&sound).Updates(models.Sound{Title: "changed title"})

	var findSound2 models.Sound
	db.First(&findSound2, findSound.ID)
	printJson("find by id sound after update", findSound2)

	// DELETE
	db.Delete(&sound)
}

func main() {
	db := dbConnect()
	db.AutoMigrate(&models.Sound{})
	dbCrudTest(db)

	app := fiber.New()

	// Middleware
	app.Use(fiberLogger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// https://apple.stackexchange.com/questions/393715/do-you-want-the-application-main-to-accept-incoming-network-connections-pop
	app.Listen("localhost:8080")
}
