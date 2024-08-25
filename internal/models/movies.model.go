package models

import (
	"encoding/json"
	"time"
)

var schemaMovies = `
CREATE TABLE public.movies (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	title varchar(255) NULL,
	image text NULL,
	director varchar(50) NULL,
	casts text NULL,
	duration varchar(20) NULL,
	release_date date NULL,
	synopsis text NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	is_deleted bool NULL DEFAULT false,
	CONSTRAINT movies_pkey PRIMARY KEY (id),
	CONSTRAINT unique_title UNIQUE (title)
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
	Is_deleted   bool      `db:"is_deleted" json:"is_deleted"`
	Created_at   time.Time `db:"created_at" json:"created_at"`
	Updated_at   time.Time `db:"updated_at" json:"updated_at"`
}

type AllMovies struct {
	Id           string     `db:"id" json:"id"`
	Title        string     `db:"title" json:"title"`
	Image        string     `db:"image" json:"image"`
	Genres       string     `db:"genres" json:"genres"`
	Director     string     `db:"director" json:"director"`
	Casts        string     `db:"casts" json:"casts"`
	Duration     string     `db:"duration" json:"duration"`
	Release_Date *time.Time `db:"release_date" json:"release_date"`
	Synopsis     string     `db:"synopsis" json:"synopsis"`
	Is_deleted   bool       `db:"is_deleted" json:"is_deleted"`
	Created_at   time.Time  `db:"created_at" json:"created_at"`
	Updated_at   time.Time  `db:"updated_at" json:"updated_at"`
}

func (u AllMovies) MarshalJSON() ([]byte, error) {
	type Alias AllMovies
	return json.Marshal(&struct {
		*Alias
		Release_Date string `json:"release_date"`
	}{
		Release_Date: formatDate(u.Release_Date),
		Alias:        (*Alias)(&u),
	})
}

func formatDate(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02")
}

type MovieResponse []AllMovies

type MoviesQuery struct {
	Page   int     `form:"page"`
	Limit  int     `form:"limit"`
	Search *string `form:"name"`
	Filter *string `form:"category"`
}

type MoviesBody struct {
	Title        string   `db:"title" json:"title" form:"title" valid:"stringlength(2|100)~Nama Movies minimal 2 dan maksimal 100"`
	Image        string   `db:"image" json:"image" valid:"-"`
	Genres       string   `json:"genres" form:"genres"`
	Director     string   `db:"director" json:"director" form:"director" valid:"stringlength(5|100)~Nama Product minimal 5 dan maksimal 100"`
	Casts        string   `db:"casts" json:"casts" form:"casts" valid:"-"`
	Duration     string   `db:"duration" json:"duration" form:"duration" valid:"-"`
	Release_Date string   `db:"release_date" json:"release_date" form:"release_date" valid:"-"`
	Synopsis     string   `db:"synopsis" json:"synopsis" form:"synopsis" valid:"-"`
	AiringDate   []string `json:"airing_date" form:"airing_date" valid:"-"`
	AiringTime   string   `json:"airing_time" form:"airing_time"`
	Locations    string   `json:"locations" form:"locations"`
}
