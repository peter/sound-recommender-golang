package initializers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/peter/sound-recommender-golang/models"
	"github.com/peter/sound-recommender-golang/util"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func DatabaseCrudTest(db *gorm.DB) {
	startTime := time.Now()
	log.Info("DatabaseCrudTest starting")
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
	util.PrintJson("created sound", sound)

	// FIND BY ID
	var findSound models.Sound
	db.First(&findSound, sound.ID)
	util.PrintJson("find by id sound", findSound)

	// UPDATE
	// https://gorm.io/docs/update.html
	// db.Model(&sound).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	db.Model(&sound).Updates(models.Sound{Title: "changed title"})

	var findSound2 models.Sound
	db.First(&findSound2, findSound.ID)
	util.PrintJson("find by id sound after update", findSound2)

	// DELETE
	db.Delete(&sound)

	elapsed := time.Since(startTime)
	log.Infof("DatabaseCrudTest done elapsed=%s", elapsed)
}
