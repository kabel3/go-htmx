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
		"Films":          films,
		"Genres":         database.GetGenres(),
		"Count":          len(films),
		"FavoriteGenres": database.GetMostPopularGenres(),
	})
}

func HandleFilmEdit(c *gin.Context) {
	filmId, _ := strconv.Atoi(c.Request.URL.Query().Get("id"))

	film := database.GetFilm(uint(filmId))
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

func GetFavoriteGenres(c *gin.Context) {
	c.HTML(http.StatusOK, "favorite-genres", map[string]interface{}{
		"FavoriteGenres": database.GetMostPopularGenres(),
	})
}

func GetFilmCount(c *gin.Context) {
	films := database.GetFilms()

	c.HTML(http.StatusOK, "film-count", map[string]interface{}{
		"Count": len(films),
	})
}

func SearchFilms(c *gin.Context) {
	keyword := c.Request.URL.Query().Get("keyword")
	films := database.SearchFilms(keyword)

	c.HTML(http.StatusOK, "film-list", map[string]interface{}{
		"Films": films,
	})
}

func AddFilm(c *gin.Context) {
	title := c.PostForm("title")
	director := c.PostForm("director")
	genre := c.PostForm("genre")

	if title != "" && director != "" {
		genreId, _ := strconv.Atoi(genre)
		film := database.AddFilm(title, director, uint(genreId))

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

		if updated := database.UpdateFilm(uint(filmId), title, director, uint(genreId)); updated {
			c.Header("HX-Location", "/")
		}
	}
}

func DeleteFilm(c *gin.Context) {
	filmId, _ := strconv.Atoi(c.Request.URL.Query().Get("id"))

	if err := database.RemoveFilm(uint(filmId)); err == nil {
		c.Header("HX-Trigger", "films-changed")
	}

	films := database.GetFilms()

	c.HTML(http.StatusOK, "film-list", map[string]interface{}{
		"Films": films,
	})
}

func StarFilm(c *gin.Context) {
	filmId, _ := strconv.Atoi(c.Request.URL.Query().Get("id"))
	starred, _ := strconv.ParseBool(c.Request.URL.Query().Get("starred"))

	if updated := database.ToggleStarredFilm(uint(filmId)); updated {
		c.Header("HX-Trigger", "films-changed")

		c.HTML(http.StatusOK, "film-starred", gin.H{
			"Id":      filmId,
			"Starred": starred,
		})
	}
}
