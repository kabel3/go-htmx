package database

import (
	"fmt"
	"kabel/packages/structs"

	_ "modernc.org/sqlite"
)

func GetFilms() ([]structs.Film, error) {
	if err := OpenDatabase(); err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT f.id, title, director, g.description FROM films f INNER JOIN genres g ON g.id = f.genreId")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var films []structs.Film

	for rows.Next() {
		var film structs.Film

		if err := rows.Scan(&film.Id, &film.Title, &film.Director, &film.Genre); err != nil {
			return films, err
		}

		films = append(films, film)
	}

	return films, nil
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
