package models

var schemaUser = `
CREATE TABLE public.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(20),
    last_name VARCHAR(20),
    email VARCHAR(50),
    password VARCHAR(255),
    role VARCHAR,
    image TEXT,
    phone_number VARCHAR(15),
    point VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`