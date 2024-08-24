package models

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