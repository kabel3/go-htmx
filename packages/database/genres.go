package database

import (
	"fmt"
	"kabel/packages/structs"

	_ "modernc.org/sqlite"
)

func SeedDefaultGenres() error {
	if err := OpenDatabase(); err != nil {
		return err
	}

	_, err := db.Exec(`
		INSERT OR IGNORE INTO genres (description) VALUES ('Action');
		INSERT OR IGNORE INTO genres (description) VALUES ('Romance');
		INSERT OR IGNORE INTO genres (description) VALUES ('Crime');
		INSERT OR IGNORE INTO genres (description) VALUES ('Suspense');
		INSERT OR IGNORE INTO genres (description) VALUES ('Aventure');
		INSERT OR IGNORE INTO genres (description) VALUES ('Fantaisie');
		INSERT OR IGNORE INTO genres (description) VALUES ('Science-fiction');	
		INSERT OR IGNORE INTO genres (description) VALUES ('Horreur');	
	`)

	return err
}

func GetGenres() ([]structs.Genre, error) {
	if err := OpenDatabase(); err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, description FROM genres")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var genres []structs.Genre

	for rows.Next() {
		var genre structs.Genre

		if err := rows.Scan(&genre.Id, &genre.Description); err != nil {
			return genres, err
		}

		genres = append(genres, genre)
	}

	if err = rows.Err(); err != nil {
		return genres, err
	}

	return genres, nil
}

func GetGenre(genreId int) structs.Genre {
	var genre structs.Genre

	if err := OpenDatabase(); err != nil {
		fmt.Println(err.Error())
		return structs.Genre{}
	}

	stmt, err := db.Prepare("SELECT id, description FROM genres WHERE id = ?")

	if err != nil {
		fmt.Println(err.Error())
		return structs.Genre{}
	}

	row := stmt.QueryRow(genreId)
	stmt.Close()

	if err := row.Scan(&genre.Id, &genre.Description); err != nil {
		fmt.Println(err.Error())
		return structs.Genre{}
	}

	return genre
}
