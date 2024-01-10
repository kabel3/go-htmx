package database

import (
	"fmt"
	"kabel/packages/database/models"
	"kabel/packages/structs"
)

func GetFilm(filmId uint) structs.Film {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return structs.Film{}
	}

	var film models.Film
	result := db.First(&film, filmId)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return structs.Film{}
	}

	return structs.Film{
		Id:       film.ID,
		Title:    film.Title,
		Director: film.Director,
		GenreId:  film.GenreId,
	}
}

func GetFilms() []structs.Film {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return []structs.Film{}
	}

	var films []structs.Film
	err := db.Model(&models.Film{}).Select(`
		films.id as Id,
		films.title as Title,
		films.director as Director,
		films.genre_id as GenreId,
		genres.description as Genre,
		films.starred as Starred
	`).Joins("INNER JOIN genres ON genres.id = films.genre_id").Scan(&films).Error

	if err != nil {
		panic(err)
	}

	return films
}

func AddFilm(title string, director string, genreId uint) structs.Film {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return structs.Film{}
	}

	if title == "" || director == "" {
		fmt.Println("Paramètres manquants pour l'ajout de film")
		return structs.Film{}
	}

	film := models.Film{Title: title, Director: director, GenreId: genreId}
	result := db.Create(&film)

	if result.Error != nil {
		panic(result.Error.Error())
	}

	var genre models.Genre
	db.First(&genre, genreId)

	return structs.Film{
		Id:       film.ID,
		Title:    film.Title,
		Director: film.Director,
		Genre:    genre.Description,
		GenreId:  genre.ID,
	}
}

func UpdateFilm(filmId uint, title string, director string, genreId uint) bool {
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

func ToggleStarredFilm(filmId uint) bool {
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

func RemoveFilm(filmId uint) error {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	db.Delete(&models.Film{}, filmId)

	return nil
}

func GetMostPopularGenres() string {
	if err := OpenDatabase(); err != nil {
		panic(err.Error())
	}

	var genre string

	subQuery := db.Model(&models.Film{}).Select(`
		films.genre_id as GenreId,
		genres.description as Description,
		COUNT(1) as GenreCount
	`).Joins(`
		INNER JOIN genres ON genres.id = films.genre_id
	`).Group("films.genre_id").Where("films.starred = 1")

	var maxCount int

	db.Table("(?)", subQuery).Select("COALESCE(MAX(GenreCount), 0)").Scan(&maxCount)

	if maxCount == 0 {
		return "-"
	}

	db.Table("(?)", subQuery).Select("GROUP_CONCAT(Description, ', ')").Where("GenreCount = ?", maxCount).Scan(&genre)

	return genre
}
