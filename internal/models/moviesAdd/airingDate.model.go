package moviesAdd

import "time"

var schemaAiringDate = `
CREATE TABLE public.airing_date (
	id serial4 NOT NULL,
	start_date date NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	end_date date NULL,
	CONSTRAINT airing_date_pkey PRIMARY KEY (id)
);
`

type AiringDate struct {
	Id         int       `db:"id" json:"id"`
	Start_date string    `db:"start_date" json:"start_date"`
	End_date   string    `db:"end_date" json:"end_date"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
