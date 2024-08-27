package moviesAdd

import "time"

var schemaAiringTimeDate = `
CREATE TABLE public.airing_time_date (
	id serial4 NOT NULL,
	airing_time_id int4 NULL,
	date_id int4 NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT airing_time_date_pkey PRIMARY KEY (id),
	CONSTRAINT airing_time_date_airing_time_id_fkey FOREIGN KEY (airing_time_id) REFERENCES public.airing_time(id) ON DELETE CASCADE,
	CONSTRAINT airing_time_date_date_id_fkey FOREIGN KEY (date_id) REFERENCES public.airing_date(id) ON DELETE CASCADE
);
`

type AiringTimeDate struct {
	Id             int       `db:"id" json:"id"`
	Airing_time_id int       `db:"airing_time_id" json:"airing_time_id"`
	Date_id        int       `db:"date_id" json:"date_id"`
	Created_at     time.Time `db:"created_at" json:"created_at"`
	Updated_at     time.Time `db:"updated_at" json:"updated_at"`
}
