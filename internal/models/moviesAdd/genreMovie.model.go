package moviesAdd

import "time"

var schemaGenreMovies = `
CREATE TABLE public.genre_movies (
	id serial4 NOT NULL,
	genre_id int4 NULL,
	movie_id uuid NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT genre_movies_pkey PRIMARY KEY (id),
	CONSTRAINT genre_movies_genre_id_fkey FOREIGN KEY (genre_id) REFERENCES public.genres(id) ON DELETE CASCADE,
	CONSTRAINT genre_movies_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE
);
`

type GenreMovie struct {
	ID         int        `db:"id" json:"id"`
	Genre_id   int        `db:"genre_id" json:"genre_id"`
	Movie_id   string     `db:"movie_id" json:"movie_id"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
