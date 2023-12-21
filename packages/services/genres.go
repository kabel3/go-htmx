package services

import (
	"fmt"
	"kabel/packages/database"
	"kabel/packages/structs"
)

func GetGenres() []structs.Genre {
	var err error
	var genres []structs.Genre

	if genres, err = database.GetGenres(); err != nil {
		fmt.Printf("Erreur à la requête de genres: %s", err.Error())
		return nil
	}

	return genres
}
