package moviesAdd

import "time"

var schemaGenreMovies = `
CREATE TABLE public.genre_movies (
    id SERIAL PRIMARY KEY,
    genre_id INT,
    movie_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (genre_id) REFERENCES genres(id),
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);
`

type GenreMovie struct {
	ID         int        `db:"id" json:"id"`
	Genre_id   int        `db:"genre_id" json:"genre_id"`
	Movie_id   string     `db:"movie_id" json:"movie_id"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
