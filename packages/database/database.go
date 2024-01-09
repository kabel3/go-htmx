package database

import (
	"kabel/packages/database/models"

	_ "modernc.org/sqlite"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var databaseURL = "./data/db.sqlite"
var db *gorm.DB

func OpenDatabase() error {
	var err error

	if db == nil {
		db, err = gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})

		if err != nil {
			return err
		}
	}

	return nil
}

func InitDatabase() error {
	if err := OpenDatabase(); err != nil {
		return err
	}

	// Migrations
	db.AutoMigrate(&models.Film{}, &models.Genre{})

	return nil
}
