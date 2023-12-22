package controllers

import (
	"kabel/packages/database"
	"kabel/packages/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleFilmIndex(c *gin.Context) {
	films := database.GetFilms()

	c.HTML(http.StatusOK, "films_index.html", map[string]interface{}{
		"Films":  films,
		"Genres": database.GetGenres(),
		"Count":  len(films),
	})
}

func HandleFilmEdit(c *gin.Context) {
	filmId, _ := strconv.Atoi(c.Request.URL.Query().Get("id"))

	film := database.GetFilm(filmId)
	genres := database.GetGenres()

	if film == (structs.Film{}) {
		c.Redirect(http.StatusNotFound, "/")
	} else {
		c.HTML(http.StatusOK, "films_edit.html", map[string]interface{}{
			"Film":   film,
			"Genres": genres,
		})
	}
}

func GetFilmCount(c *gin.Context) {
	films := database.GetFilms()

	c.HTML(http.StatusOK, "film-count", map[string]interface{}{
		"Count": len(films),
	})
}

func AddFilm(c *gin.Context) {
	title := c.PostForm("title")
	director := c.PostForm("director")
	genre := c.PostForm("genre")

	if title != "" && director != "" {
		genreId, _ := strconv.Atoi(genre)
		film := database.AddFilm(title, director, genreId)

		if film != (structs.Film{}) {
			c.Header("HX-Trigger", "films-changed")
		}

		c.HTML(http.StatusOK, "film-list-element", film)
	}
}

func UpdateFilm(c *gin.Context) {
	id := c.PostForm("id")
	title := c.PostForm("title")
	director := c.PostForm("director")
	genre := c.PostForm("genre")

	filmId, _ := strconv.Atoi(id)

	if title != "" && director != "" {
		genreId, _ := strconv.Atoi(genre)
		updated := database.UpdateFilm(filmId, title, director, genreId)

		if updated {
			c.Header("HX-Location", "/")
		}
	}
}

func DeleteFilm(c *gin.Context) {
	filmId, _ := strconv.Atoi(c.Request.URL.Query().Get("id"))

	if err := database.RemoveFilm(filmId); err == nil {
		c.Header("HX-Trigger", "films-changed")
	}

	films := database.GetFilms()

	c.HTML(http.StatusOK, "film-list", map[string]interface{}{
		"Films": films,
	})
}
