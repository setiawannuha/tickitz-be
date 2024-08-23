package models

var schemaAiringTime = `
CREATE TABLE public.airing_time (
    id SERIAL PRIMARY KEY,
    time TIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`