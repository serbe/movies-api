package main

import "github.com/go-pg/pg"

var db *pg.DB

func initDB() {
	db = pg.Connect(&pg.Options{
		Database: "movies",
		User:     "movuser",
		Password: "movpass",
	})
}

func getMovieByID(id int64) (Movie, error) {
	var movie Movie
	err := db.Model(&movie).Where("id = ?", id).Select()
	if err != nil {
		return movie, err
	}
	movie.Torrents, err = getTottentsByMovieID(id)
	return movie, err
}

func getAllMovies(offset int) ([]Movie, error) {
	var movies []Movie
	err := db.Model(&movies).Offset(offset).Limit(100).Select()
	if err != nil {
		return movies, err
	}
	for i := range movies {
		movies[i].Torrents, _ = getTottentsByMovieID(movies[i].ID)
	}
	return movies, err
}

func getAllMoviesByYear(year int) ([]Movie, error) {
	var movies []Movie
	err := db.Model(&movies).Where("year = ?", year).Limit(100).Select()
	if err != nil {
		return movies, err
	}
	for i := range movies {
		movies[i].Torrents, _ = getTottentsByMovieID(movies[i].ID)
	}
	return movies, err
}

func getTorrentByID(id int64) (Torrent, error) {
	var torrent Torrent
	err := db.Model(&torrent).Where("id = ?", id).Select()
	return torrent, err
}

func getTottentsByMovieID(id int64) ([]Torrent, error) {
	var torrents []Torrent
	err := db.Model(&torrents).Where("movie_id = ?", id).Select()
	return torrents, err
}

func getAllTorrents() ([]Torrent, error) {
	var torrents []Torrent
	err := db.Select(&torrents)
	return torrents, err
}
