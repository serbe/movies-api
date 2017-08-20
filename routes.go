package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func getMovie(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Status string `json:"status"`
		Title  string `json:"title"`
		Movie  Movie  `json:"movie"`
	}
	id := toInt64(chi.URLParam(r, "id"))
	movie, err := getMovieByID(id)
	if err != nil {
		log.Println("getMovie getMovieByID", err)
		renderError(w, r, "error in get movie")
		return
	}
	ctx := context{Status: "ok", Title: "Movie", Movie: movie}
	render.JSON(w, r, ctx)
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Status string  `json:"status"`
		Title  string  `json:"title"`
		Movies []Movie `json:"movies"`
	}
	offset := toInt(chi.URLParam(r, "offset"))
	movies, err := getAllMovies(offset)
	if err != nil {
		log.Println("listMovies getAllMovies", err)
		renderError(w, r, "error in get all movies")
		return
	}
	ctx := context{Status: "ok", Title: "Movies list", Movies: movies}
	render.JSON(w, r, ctx)
}

func listMoviesByYear(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Status string  `json:"status"`
		Title  string  `json:"title"`
		Movies []Movie `json:"movies"`
	}
	year := toInt(chi.URLParam(r, "year"))
	movies, err := getAllMoviesByYear(year)
	if err != nil {
		log.Println("listMoviesByYear getAllMoviesByYear", err)
		renderError(w, r, "error in get all movies")
		return
	}
	ctx := context{Status: "ok", Title: "Movies list", Movies: movies}
	render.JSON(w, r, ctx)
}

func listTorrents(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Status   string    `json:"status"`
		Title    string    `json:"title"`
		Torrents []Torrent `json:"torrents"`
	}
	torrents, err := getAllTorrents()
	if err != nil {
		log.Println("listTorrents getAllTorrents", err)
		renderError(w, r, "error in get all torrents")
		return
	}
	ctx := context{Status: "ok", Title: "Movies list", Torrents: torrents}
	render.JSON(w, r, ctx)
}

func getTorrent(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Status  string  `json:"status"`
		Title   string  `json:"title"`
		Torrent Torrent `json:"torrent"`
	}
	id := toInt64(chi.URLParam(r, "id"))
	torrent, err := getTorrentByID(id)
	if err != nil {
		log.Println("getTorrent getTorrentByID", err)
		renderError(w, r, "error in get torrent")
		return
	}
	ctx := context{Status: "ok", Title: "Torrent", Torrent: torrent}
	render.JSON(w, r, ctx)
}
