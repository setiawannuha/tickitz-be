package moviesAdd

import "time"

var schemaLocations = `
CREATE TABLE public.locations (
	id serial4 NOT NULL,
	"name" varchar(50) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT locations_pkey PRIMARY KEY (id)
);
`

type Locations struct {
	ID         int        `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
}
