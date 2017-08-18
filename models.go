package main

// Movie all values
// TableName     Название таблицы
// ID            id
// Section       Раздел форума
// Name          Название
// EngName       Английское название
// Year          Год
// Genre         Жанр
// Country       Производство
// Director      Режиссер
// Producer      Продюсер
// Actors        Актеры
// Description   Описание
// Age           Возраст
// ReleaseDate   Дата мировой премьеры
// RussianDate   Дата премьеры в России
// Duration      Продолжительность
// Kinopoisk     Рейтинг кинопоиска
// Imdb          Рейтинг IMDb
// Poster        Имя файла постера
// PosterURL     Сетевая ссылка на постер
type Movie struct {
	ID          int64     `sql:"id"                        json:"id"`
	Section     string    `sql:"section"                   json:"section"`
	Name        string    `sql:"name"                      json:"name"`
	EngName     string    `sql:"eng_name"                  json:"eng_name"`
	Year        int       `sql:"year"                      json:"year"`
	Genre       []string  `sql:"genre"        pg:",array"  json:"genre"`
	Country     []string  `sql:"country"      pg:",array"  json:"country"`
	RawCountry  string    `sql:"raw_country"               json:"-"`
	Director    []string  `sql:"director"     pg:",array"  json:"director"`
	Producer    []string  `sql:"producer"     pg:",array"  json:"producer"`
	Actor       []string  `sql:"actor"        pg:",array"  json:"actor"`
	Description string    `sql:"description"               json:"description"`
	Age         string    `sql:"age"                       json:"age"`
	ReleaseDate string    `sql:"release_date"              json:"release_date"`
	RussianDate string    `sql:"russian_date"              json:"russian_date"`
	Duration    string    `sql:"duration"                  json:"duration"`
	Kinopoisk   float64   `sql:"kinopoisk"                 json:"kinopoisk"`
	IMDb        float64   `sql:"imdb"                      json:"imdb"`
	Poster      string    `sql:"poster"                    json:"poster"`
	PosterURL   string    `sql:"poster_url"                json:"poster_url"`
	Torrents    []Torrent `sql:"-"                         json:"torrents"`
}

// Torrent all values
// TableName     Название таблицы
// ID            id
// FilmID        Указатель на фильм
// DateCreate    Дата создания раздачи
// Href          Ссылка
// Torrent       Ссылка на torrent
// Magnet        Ссылка на magnet
// NNM           Рейтинг nnm-club
// SubtitlesType Вид субтитров
// Subtitles     Субтитры
// Video         Видео
// Quality       Качество видео
// Resolution    Разрешение видео
// Audio1        Аудио1
// Audio2        Аудио2
// Audio3        Аудио3
// Translation   Перевод
// Size          Размер
// Seeders       Количество раздающих
// Leechers      Количество скачивающих
type Torrent struct {
	ID            int64   `sql:"id"`
	MovieID       int64   `sql:"movie_id"`
	DateCreate    string  `sql:"date_create"`
	Href          string  `sql:"href"`
	Torrent       string  `sql:"torrent"`
	Magnet        string  `sql:"magnet"`
	NNM           float64 `sql:"nnm"`
	SubtitlesType string  `sql:"subtitles_type"`
	Subtitles     string  `sql:"subtitles"`
	Video         string  `sql:"video"`
	Quality       string  `sql:"quality"`
	Resolution    string  `sql:"resolution"`
	Audio1        string  `sql:"audio1"`
	Audio2        string  `sql:"audio2"`
	Audio3        string  `sql:"audio3"`
	Translation   string  `sql:"translation"`
	Size          int     `sql:"size"`
	Seeders       int     `sql:"seeders"`
	Leechers      int     `sql:"leechers"`
}
