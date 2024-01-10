package main

import (
	"fmt"
	"kabel/packages/controllers"
	"kabel/packages/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("En écoute sur localhost:8080...")

	err := database.InitDatabase()

	if err != nil {
		fmt.Println("Erreur initialisation DB " + err.Error())
		os.Exit(1)
	}

	err = database.SeedDefaultGenres()

	if err != nil {
		fmt.Println("Erreur initialisation des genres " + err.Error())
		os.Exit(1)
	}

	r := gin.Default()

	r.LoadHTMLGlob("./templates/**/*")
	r.Static("/static/", "./static")

	r.GET("/", controllers.HandleFilmIndex)
	r.GET("/film", controllers.HandleFilmEdit)

	r.GET("/api/films/count", controllers.GetFilmCount)
	r.GET("/api/films/favorite-genres", controllers.GetFavoriteGenres)
	r.GET("/api/films/search", controllers.SearchFilms)

	r.PUT("/api/film", controllers.AddFilm)
	r.POST("/api/film", controllers.UpdateFilm)
	r.DELETE("/api/film", controllers.DeleteFilm)

	r.POST("/api/film/star", controllers.StarFilm)

	log.Fatal(r.Run("localhost:8080"))
}
