package moviesAdd

import "time"

var schemaGenres = `
CREATE TABLE public.genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

type Genres struct {
	ID         int        `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
}
