package controllers

import (
	"kabel/packages/database"
	"kabel/packages/database/models"
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

	if film == (models.Film{}) {
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
		genreId, _ := strconv.ParseUint(genre, 10, 64)
		film := database.AddFilm(title, director, genreId)

		if film != (models.Film{}) {
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

	filmId, _ := strconv.ParseUint(id, 10, 64)

	if title != "" && director != "" {
		genreId, _ := strconv.ParseUint(genre, 10, 64)

		if updated := database.UpdateFilm(filmId, title, director, genreId); updated {
			c.Header("HX-Location", "/")
		}
	}
}

func DeleteFilm(c *gin.Context) {
	filmId, _ := strconv.ParseUint(c.Request.URL.Query().Get("id"), 10, 64)

	if err := database.RemoveFilm(filmId); err == nil {
		c.Header("HX-Trigger", "films-changed")
	}

	films := database.GetFilms()

	c.HTML(http.StatusOK, "film-list", map[string]interface{}{
		"Films": films,
	})
}

func StarFilm(c *gin.Context) {
	filmId, _ := strconv.ParseUint(c.Request.URL.Query().Get("id"), 10, 64)
	starred, _ := strconv.ParseBool(c.Request.URL.Query().Get("starred"))

	if updated := database.ToggleStarredFilm(filmId); updated {
		c.HTML(http.StatusOK, "film-starred", gin.H{
			"Id":      filmId,
			"Starred": starred,
		})
	}
}
