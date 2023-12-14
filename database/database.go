package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=sound-recommender port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Verbose logging of SQL queries with Info level
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
