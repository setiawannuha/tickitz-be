package models

var schemaMovies = `
CREATE TABLE public.movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255),
    image TEXT,
    director VARCHAR(50),
    casts TEXT,
    duration VARCHAR(20),
    release_date DATE,
    synopsis TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`