package initializers

import (
	"github.com/peter/sound-recommender-golang/models"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.Sound{})
}
