package main

import "github.com/go-pg/pg"

func initDB() *pg.DB {
	return pg.Connect(&pg.Options{
		Database: "movies",
		User:     "movuser",
		Password: "md5ceb6f50cdffcd1cb6c2bddf04282ee6f",
	})
}

func getMovieByID(db *pg.DB, id int64) (Movie, error) {
	movie := Movie{ID: id}
	err := db.Select(&movie)
	if err != nil {
		return movie, err
	}
	movie.Torrents, err = getTottentsByMovieID(db, id)
	return movie, err
}

func getAllMovies(db *pg.DB) ([]Movie, error) {
	var movies []Movie
	err := db.Select(&movies)
	if err != nil {
		return movies, err
	}
	for i := range movies {
		movies[i].Torrents, _ = getTottentsByMovieID(db, movies[i].ID)
	}
	return movies, err
}

func getTottentsByMovieID(db *pg.DB, id int64) ([]Torrent, error) {
	var torrents []Torrent
	err := db.Model(&torrents).Where("movie_id = ?", id).Select()
	return torrents, err
}
