package moviesAdd

import "time"

var schemaAiringTime = `
CREATE TABLE public.airing_time (
	id serial4 NOT NULL,
	"time" time NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT airing_time_pkey PRIMARY KEY (id)
);
`

type AiringTime struct {
	Id         int       `db:"id" json:"id"`
	Time       time.Time `db:"time" json:"time"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
