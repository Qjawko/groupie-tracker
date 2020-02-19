package controller

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/service"
	"html/template"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {


	id := r.URL.Path[len("/artist/"):]
	if len(id) == 0 {
		w.Header().Set("Content-Type", "application/json")
		if err := getArtists(w); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			fmt.Println(err.Error())
		}
	} else {
		idArtist, err := strconv.Atoi(id)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			fmt.Println(err.Error())
			return
		}


		artist, err := service.GetArtistById(idArtist)

		t, err := template.ParseFiles("templates/base.html", "templates/artist.html")
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		err = t.Execute(w, artist)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			fmt.Println(err.Error())
			return
		}
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, err := template.ParseFiles("templates/base.html", "templates/home.html")

		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}
		artists, err := service.Get()
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}
		err = t.Execute(w, artists)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}
}

func getArtists(w http.ResponseWriter) error {
	artists, err := service.Get()
	if err != nil {
		return err
	}

	artistsJson, err := json.Marshal(artists)
	if err != nil {
		return err
	}

	_, err = w.Write(artistsJson)
	if err != nil {
		return err
	}

	return nil
}
