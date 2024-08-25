package moviesAdd

import "time"

var schemaAiringTime = `
CREATE TABLE public.airing_time (
    id SERIAL PRIMARY KEY,
    time TIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

type AiringTime struct {
	Id         int       `db:"id" json:"id"`
	Time       time.Time `db:"time" json:"time"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at"`
}
