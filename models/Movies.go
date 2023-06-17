package models

import "time"

type Movie struct {
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	Budget           int     `json:"budget"`
	Genres           []Genre `json:"genres" gorm:"many2many:movie_genres"`
	Homepage         string  `json:"homepage"`
	ID               int     `json:"id" gorm:"primarykey"`
	ImdbID           string  `json:"imdb_id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Revenue          int     `json:"revenue"`
	Runtime          int     `json:"runtime"`
	Status           string  `json:"status"`
	Tagline          string  `json:"tagline"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	ListID           uint    `json:"list"`
	Watched          bool    `json:"watched"`
}
