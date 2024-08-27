package moviesAdd

import "time"

var schemaSeats = `
CREATE TABLE public.seats (
	id serial4 NOT NULL,
	"name" varchar(10) NULL,
	status varchar(10) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT seats_pkey PRIMARY KEY (id)
);
`

type Seats struct {
	ID         int        `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Status     string     `db:"status" json:"status"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
}
