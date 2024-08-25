package moviesAdd

import "time"

var schemaLocationsMovieTimes = `
CREATE TABLE public.location_movie_time (
    id SERIAL PRIMARY KEY,
    location_id INT,
    movie_time_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (location_id) REFERENCES locations(id),
    FOREIGN KEY (movie_time_id) REFERENCES movies_time(id)
);
`

type LocationMovieTime struct {
	ID            int        `db:"id" json:"id"`
	Location_id   int        `db:"location_id" json:"location_id"`
	Movie_time_id int        `db:"movie_time_id" json:"movie_time_id"`
	Created_at    *time.Time `db:"created_at" json:"created_at"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at"`
}
