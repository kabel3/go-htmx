package main

import (
	"fmt"
	"html/template"
	"kabel/packages/database"
	"kabel/packages/services"
	"kabel/packages/structs"
	"log"
	"net/http"
	"os"
	"strconv"
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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(
			"./pages/_base.html",
			"./pages/films.html",
		))

		films := services.GetFilms()
		genres := services.GetGenres()

		filmList := map[string]interface{}{
			"Films":  services.GetFilms(),
			"Genres": genres,
			"Count":  len(films),
		}

		tmpl.ExecuteTemplate(w, "base", filmList)
	})

	http.HandleFunc("/film-count", func(w http.ResponseWriter, r *http.Request) {
		films := services.GetFilms()

		tmpl := template.Must(template.ParseFiles("./ui/html/pages/films.html"))
		tmpl.ExecuteTemplate(w, "film-count", map[string]interface{}{
			"Count": len(films),
		})
	})

	http.HandleFunc("/add-film/", func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		genre := r.PostFormValue("genre")

		if title != "" && director != "" {
			genreId, _ := strconv.Atoi(genre)
			film := database.AddFilm(title, director, genreId)

			if film != (structs.Film{}) {
				w.Header().Set("HX-Trigger", "films-changed")
			}

			tmpl := template.Must(template.ParseFiles("./ui/html/pages/films.html"))
			tmpl.ExecuteTemplate(w, "film-list-element", film)
		}
	})

	http.HandleFunc("/remove-film", func(w http.ResponseWriter, r *http.Request) {
		filmId, _ := strconv.Atoi(r.URL.Query().Get("id"))

		if err := database.RemoveFilm(filmId); err == nil {
			w.Header().Set("HX-Trigger", "films-changed")
		}

		films := services.GetFilms()

		filmList := map[string]interface{}{
			"Films": films,
		}

		tmpl := template.Must(template.ParseFiles("./ui/html/pages/films.html"))
		tmpl.ExecuteTemplate(w, "film-list", filmList)
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
