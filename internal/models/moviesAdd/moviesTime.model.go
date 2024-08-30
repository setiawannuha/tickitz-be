package moviesAdd

import "time"

var schemaMoviesTime = `
CREATE TABLE public.movies_time (
	id serial4 NOT NULL,
	movie_id uuid NULL,
	airing_time_date_id int4 NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT movies_time_pkey PRIMARY KEY (id),
	CONSTRAINT movies_time_airing_time_date_id_fkey FOREIGN KEY (airing_time_date_id) REFERENCES public.airing_time_date(id) ON DELETE CASCADE,
	CONSTRAINT movies_time_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE
);
`

type MovieTime struct {
	ID                  int        `db:"id" json:"id"`
	Movie_id            string     `db:"movie_id" json:"movie_id"`
	Airing_time_date_id int        `db:"airing_time_date_id" json:"airing_time_date_id"`
	Created_at          *time.Time `db:"created_at" json:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at"`
}
