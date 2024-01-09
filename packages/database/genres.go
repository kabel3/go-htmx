package database

import (
	"fmt"
	"kabel/packages/database/models"

	"errors"

	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func SeedDefaultGenres() error {
	var err error

	if err = OpenDatabase(); err != nil {
		return err
	}

	if err = db.AutoMigrate(&models.Genre{}); err == nil && db.Migrator().HasTable(&models.Genre{}) {
		if err := db.First(&models.Genre{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&models.Genre{Description: "Action"})
			db.Create(&models.Genre{Description: "Romance"})
			db.Create(&models.Genre{Description: "Crime"})
			db.Create(&models.Genre{Description: "Suspense"})
			db.Create(&models.Genre{Description: "Aventure"})
			db.Create(&models.Genre{Description: "Fantaisie"})
			db.Create(&models.Genre{Description: "Science-fiction"})
			db.Create(&models.Genre{Description: "Horreur"})
		}
	}

	return nil
}

func GetGenres() []models.Genre {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return []models.Genre{}
	}

	var genres []models.Genre
	result := db.Find(&genres)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return []models.Genre{}
	}

	return genres
}

func GetGenre(genreId int) models.Genre {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return models.Genre{}
	}

	var genre models.Genre
	result := db.First(&genre, genreId)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return models.Genre{}
	}

	return genre
}
