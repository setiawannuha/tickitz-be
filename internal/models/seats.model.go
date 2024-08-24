package models

var schemaSeats = `
CREATE TABLE public.seats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(10),
    status VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`