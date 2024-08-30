package models

import "time"

var schemaPayments = `
CREATE TABLE public.payment_methods (
	id serial4 NOT NULL,
	"name" varchar NULL,
	image text NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT payment_methods_pkey PRIMARY KEY (id)
);
`

type Payments struct {
	ID         int        `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Image      string     `db:"image" json:"image"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
}
