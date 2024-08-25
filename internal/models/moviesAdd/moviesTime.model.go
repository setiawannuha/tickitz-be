package moviesAdd

import "time"

var schemaMoviesTime = `
CREATE TABLE public.movies_time (
    id SERIAL PRIMARY KEY,
    movie_id UUID,
    airing_time_date_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (airing_time_date_id) REFERENCES airing_time_date(id)
);
`

type MovieTime struct {
	ID                  int        `db:"id" json:"id"`
	Movie_id            string     `db:"movie_id" json:"movie_id"`
	Airing_time_date_id int        `db:"airing_time_date_id" json:"airing_time_date_id"`
	Created_at          *time.Time `db:"created_at" json:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at"`
}
