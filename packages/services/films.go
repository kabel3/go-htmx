package services

import (
	"fmt"
	"kabel/packages/database"
	"kabel/packages/structs"
)

func GetFilms() []structs.Film {
	var err error
	var films []structs.Film

	if films, err = database.GetFilms(); err != nil {
		fmt.Printf("Erreur à la requête de films: %s", err.Error())
		return []structs.Film{}
	}

	return films
}
