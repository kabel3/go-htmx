package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var databaseURL = "./data/db.sqlite"
var db *sql.DB

func OpenDatabase() error {
	var err error

	if db == nil {
		if db, err = sql.Open("sqlite", databaseURL); err != nil {
			return err
		}
	}

	return nil
}

func InitDatabase() error {
	if err := OpenDatabase(); err != nil {
		return err
	}

	// Création de tables
	createMoviesTable := `CREATE TABLE IF NOT EXISTS films (
		id 				INTEGER PRIMARY KEY,
		title 		TEXT NOT NULL,
		director	TEXT NOT NULL,
		genreId 	INTEGER NOT NULL,
		starred		INTEGER NOT NULL DEFAULT 0 CHECK(starred IN (0,1))
	)`

	createGenresTable := `CREATE TABLE IF NOT EXISTS genres (
		id						INTEGER PRIMARY KEY,
		description		TEXT NOT NULL,
		UNIQUE(description)
	)`

	_, err := db.Exec(createMoviesTable)
	_, err = db.Exec(createGenresTable)

	return err
}
