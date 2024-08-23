package models

var schemaAiringTimeDate = `
CREATE TABLE public.airing_time_date (
    id SERIAL PRIMARY KEY,
    airing_time_id INT,
    date_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (airing_time_id) REFERENCES airing_time(id),
    FOREIGN KEY (date_id) REFERENCES airing_date(id)
);
`