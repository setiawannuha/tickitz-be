package moviesAdd

import "time"

var schemaGenres = `
CREATE TABLE public.genres (
	id serial4 NOT NULL,
	"name" varchar(20) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT genres_pkey PRIMARY KEY (id)
);
`

type Genres struct {
	ID         int        `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
}
