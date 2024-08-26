package moviesAdd

import "time"

var schemaAiringDate = `
CREATE TABLE public.airing_date (
    id SERIAL PRIMARY KEY,
    start_date DATE,
		end_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

type AiringDate struct {
	Id         int       `db:"id" json:"id"`
	Start_date string    `db:"start_date" json:"start_date"`
	End_date   string    `db:"end_date" json:"end_date"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
