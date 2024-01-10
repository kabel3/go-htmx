package database

import (
	"fmt"
	"kabel/packages/database/models"
	"kabel/packages/structs"

	"errors"

	"gorm.io/gorm"
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

func GetGenres() []structs.Genre {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return []structs.Genre{}
	}

	var genres []structs.Genre
	result := db.Model(&models.Genre{}).Scan(&genres)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return []structs.Genre{}
	}

	return genres
}

func GetGenre(genreId uint) structs.Genre {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return structs.Genre{}
	}

	var genre models.Genre
	result := db.First(&genre, genreId)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return structs.Genre{}
	}

	return structs.Genre{
		Id:          genre.ID,
		Description: genre.Description,
	}
}
