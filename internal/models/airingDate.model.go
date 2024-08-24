package models

var schemaAiringDate = `
CREATE TABLE public.airing_date (
    id SERIAL PRIMARY KEY,
    date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`