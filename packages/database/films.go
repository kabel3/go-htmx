package database

import (
	"fmt"
	"kabel/packages/database/models"
	"kabel/packages/structs"

	_ "modernc.org/sqlite"
)

func GetFilm(filmId int) models.Film {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return models.Film{}
	}

	var film models.Film
	result := db.First(&film, filmId)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return models.Film{}
	}

	return film
}

func GetFilms() []structs.Film {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return []structs.Film{}
	}

	var films []structs.Film
	err := db.Model(&models.Film{}).Select(`
		films.id as filmId,
		films.title as title,
		films.director as director,
		films.genre_id as genreId,
		genres.description as genre
	`).Joins("INNER JOIN genres ON genres.id = films.genre_id").Scan(&films).Error

	fmt.Println(films)

	if err != nil {
		panic(err)
	}

	return films
}

func AddFilm(title string, director string, genreId uint64) models.Film {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return models.Film{}
	}

	if title == "" || director == "" {
		fmt.Println("Paramètres manquants pour l'ajout de film")
		return models.Film{}
	}

	film := models.Film{Title: title, Director: director, GenreId: genreId}
	result := db.Create(&film)

	if result.Error != nil {
		panic(result.Error.Error())
	}

	return film
}

func UpdateFilm(filmId uint64, title string, director string, genreId uint64) bool {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return false
	}

	if title == "" || director == "" {
		fmt.Println("Paramètres manquants pour l'ajout de film")
		return false
	}

	var film models.Film
	db.First(&film, filmId)

	film.Title = title
	film.Director = director
	film.GenreId = genreId

	db.Save(&film)

	return true
}

func ToggleStarredFilm(filmId uint64) bool {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return false
	}

	var film models.Film
	db.First(&film, filmId)

	film.Starred = !film.Starred

	db.Save(&film)

	return true
}

func RemoveFilm(filmId uint64) error {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	db.Delete(&models.Film{}, filmId)

	return nil
}
