package models

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