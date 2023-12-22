package database

import (
	"fmt"
	"kabel/packages/structs"

	_ "modernc.org/sqlite"
)

func GetFilm(filmId int) structs.Film {
	var film structs.Film

	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return structs.Film{}
	}

	stmt, err := db.Prepare("SELECT id, title, director, genreId FROM films WHERE id = ?")

	if err != nil {
		fmt.Println(err.Error())
		return structs.Film{}
	}

	row := stmt.QueryRow(filmId)
	stmt.Close()

	if err := row.Scan(&film.Id, &film.Title, &film.Director, &film.GenreId); err != nil {
		fmt.Println(err.Error())
		return structs.Film{}
	}

	return film
}

func GetFilms() []structs.Film {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return []structs.Film{}
	}

	rows, err := db.Query("SELECT f.id, title, director, g.description FROM films f INNER JOIN genres g ON g.id = f.genreId")

	if err != nil {
		fmt.Println(err.Error())
		return []structs.Film{}
	}

	defer rows.Close()

	var films []structs.Film

	for rows.Next() {
		var film structs.Film

		if err := rows.Scan(&film.Id, &film.Title, &film.Director, &film.Genre); err != nil {
			fmt.Println(err.Error())
			return films
		}

		films = append(films, film)
	}

	return films
}

func AddFilm(title string, director string, genreId int) structs.Film {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return structs.Film{}
	}

	if title == "" || director == "" {
		fmt.Println("Paramètres manquants pour l'ajout de film")
		return structs.Film{}
	}

	stmt, err := db.Prepare("INSERT INTO films (title, director, genreId) VALUES (?, ?, ?) ")

	if err != nil {
		fmt.Println(err.Error())
		return structs.Film{}
	}

	res, err := stmt.Exec(title, director, genreId)
	stmt.Close()

	if err != nil {
		fmt.Println(err.Error())
		return structs.Film{}
	}

	lastInsertedId, _ := res.LastInsertId()

	var film structs.Film

	film.Id = int(lastInsertedId)
	film.Title = title
	film.Director = director

	genre := GetGenre(genreId)

	if genre == (structs.Genre{}) {
		fmt.Println("Genre de film non-trouvé")
		return film
	}

	film.Genre = genre.Description

	return film
}

func UpdateFilm(filmId int, title string, director string, genreId int) bool {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return false
	}

	if title == "" || director == "" {
		fmt.Println("Paramètres manquants pour l'ajout de film")
		return false
	}

	stmt, err := db.Prepare("UPDATE films SET title = ?, director = ?, genreId = ? WHERE id = ?")

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	_, err = stmt.Exec(title, director, genreId, filmId)
	stmt.Close()

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func RemoveFilm(filmId int) error {
	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	stmt, err := db.Prepare("DELETE FROM films WHERE id = ?")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(filmId)
	stmt.Close()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
