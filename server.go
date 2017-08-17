package main

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func initServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(corsHandler)

	FileServer(r, "/static", http.Dir(filepath.Join("public", "static")))

	r.Group(func(r chi.Router) {
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Route("/api/v1/movies", func(r chi.Router) {
			r.Get("/", listMovies)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", getMovie)
			})
		})

		r.Route("/api/v1/torrents", func(r chi.Router) {
			r.Get("/", listTorrents)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", getTorrent)
			})
		})
	})

	http.ListenAndServe("127.0.0.1:6060", r)
}
