package models

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