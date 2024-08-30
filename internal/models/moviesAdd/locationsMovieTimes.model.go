package moviesAdd

import "time"

var schemaLocationsMovieTimes = `
CREATE TABLE public.location_movie_time (
	id serial4 NOT NULL,
	location_id int4 NULL,
	movie_time_id int4 NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT location_movie_time_pkey PRIMARY KEY (id),
	CONSTRAINT location_movie_time_location_id_fkey FOREIGN KEY (location_id) REFERENCES public.locations(id) ON DELETE CASCADE,
	CONSTRAINT location_movie_time_movie_time_id_fkey FOREIGN KEY (movie_time_id) REFERENCES public.movies_time(id) ON DELETE CASCADE
);
`

type LocationMovieTime struct {
	ID            int        `db:"id" json:"id"`
	Location_id   int        `db:"location_id" json:"location_id"`
	Movie_time_id int        `db:"movie_time_id" json:"movie_time_id"`
	Created_at    *time.Time `db:"created_at" json:"created_at"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at"`
}
