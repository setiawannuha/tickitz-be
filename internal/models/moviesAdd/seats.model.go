package moviesAdd

import "time"

var schemaSeats = `
CREATE TABLE public.seats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(10),
    status VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

type Seats struct {
	ID         int        `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Status     string     `db:"status" json:"status"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
}
