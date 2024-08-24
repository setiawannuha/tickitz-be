package models

var schemaGenres = `
CREATE TABLE public.genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`