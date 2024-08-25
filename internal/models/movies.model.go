package models

import (
	"khalifgfrz/coffee-shop-be-go/internal/models/moviesAdd"
	"time"
)

var schemaMovies = `
CREATE TABLE public.movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255),
    image TEXT,
    director VARCHAR(50),
    casts TEXT,
    duration VARCHAR(20),
    release_date DATE,
    synopsis TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

type Movies struct {
	Id           string    `db:"id" json:"id"`
	Title        string    `db:"title" json:"title" form:"title" valid:"stringlength(2|100)~Nama Movies minimal 2 dan maksimal 100"`
	Image        string    `db:"image" json:"image" valid:"-"`
	Director     string    `db:"director" json:"director" form:"director" valid:"stringlength(5|100)~Nama Product minimal 5 dan maksimal 100"`
	Casts        string    `db:"casts" json:"casts" form:"casts" valid:"-"`
	Duration     string    `db:"duration" json:"duration" form:"duration" valid:"-"`
	Release_Date time.Time `db:"release_date" json:"release_date" form:"release_date" valid:"date"`
	Synopsis     string    `db:"synopsis" json:"synopsis" form:"synopsis" valid:"-"`
	Created_at   time.Time `db:"created_at" json:"created_at"`
	Updated_at   time.Time `db:"updated_at" json:"updated_at"`
}

type MovieResponse []Movies

type MoviesQuery struct {
	Page   int     `form:"page"`
	Limit  int     `form:"limit"`
	Search *string `form:"name"`
	Filter *string `form:"category"`
}

type MoviesBody struct {
	Id                string                        `db:"id" json:"id"`
	Title             string                        `db:"title" json:"title" form:"title" valid:"stringlength(2|100)~Nama Movies minimal 2 dan maksimal 100"`
	Image             string                        `db:"image" json:"image" valid:"-"`
	Genres            []moviesAdd.GenreMovie        `json:"genres" valid:"-"`
	Director          string                        `db:"director" json:"director" form:"director" valid:"stringlength(5|100)~Nama Product minimal 5 dan maksimal 100"`
	Casts             string                        `db:"casts" json:"casts" form:"casts" valid:"-"`
	Duration          string                        `db:"duration" json:"duration" form:"duration" valid:"-"`
	Release_Date      time.Time                     `db:"release_date" json:"release_date" form:"release_date" valid:"date"`
	Synopsis          string                        `db:"synopsis" json:"synopsis" form:"synopsis" valid:"-"`
	AiringDate        []moviesAdd.AiringDate        `json:"airing_date" valid:"-"`
	LocationMovieTime []moviesAdd.LocationMovieTime `json:"location_movie_time"`
	MovieTime         []moviesAdd.MovieTime         `json:"movie_time" valid:"-"`
	Created_at        *time.Time                    `db:"created_at" json:"created_at"`
	Updated_at        *time.Time                    `db:"updated_at" json:"updated_at"`
}
