package moviesAdd

import "time"

var schemaMoviesTime = `
CREATE TABLE public.orders (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	order_number text NULL,
	user_id uuid NULL,
	payment_method_id int4 NULL,
	seat_count int4 NULL,
	ticket_status varchar(20) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	total int4 NULL,
	movie_id uuid NULL,
	"date" date NULL,
	"time" time NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id),
	CONSTRAINT orders_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id),
	CONSTRAINT orders_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(id),
	CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id)
);
`

type MovieTime struct {
	ID                  int        `db:"id" json:"id"`
	Movie_id            string     `db:"movie_id" json:"movie_id"`
	Airing_time_date_id int        `db:"airing_time_date_id" json:"airing_time_date_id"`
	Created_at          *time.Time `db:"created_at" json:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at"`
}
