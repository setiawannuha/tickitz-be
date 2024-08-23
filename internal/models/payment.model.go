package models

var schemaPayments = `
CREATE TABLE public.payment_methods (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    image TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`