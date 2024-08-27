package models

import "time"

var schemaOrders = `
CREATE TABLE public.orders (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	order_number text NULL,
	user_id uuid NULL,
	payment_method_id int4 NULL,
	movie_time_id int4 NULL,
	seat_count int4 NULL,
	ticket_status varchar(20) NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	total int4 NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id),
	CONSTRAINT orders_movie_time_id_fkey FOREIGN KEY (movie_time_id) REFERENCES public.movies_time(id),
	CONSTRAINT orders_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(id),
	CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id)
);
`

type GetOrder struct {
	Id                string            `db:"id" json:"id" form:"id"`
	Order_number      *string           `db:"order_number" form:"order_number" json:"order_number"`
	User_id           *string           `db:"user_id" form:"user_id" json:"user_id"`
	Payment_method_id *string           `db:"payment_method_id" form:"payment_method_id" json:"payment_method_id"`
	Movie_time_id     *string           `db:"movie_time_id" form:"movie_time_id" json:"movie_time_id"`
	Seat_count        *string           `db:"seat_count" form:"seat_count" json:"seat_count"`
	Ticket_status     *string           `db:"ticket_status" form:"ticket_status" json:"ticket_status"`
	Orders            []GetOrderDetails `json:"orders"`
	Created_at        *time.Time        `db:"created_at" form:"created_at" json:"created_at"`
	Updated_at        *time.Time        `db:"updated_at" form:"updated_at" json:"updated_at"`
}

type GetOrders []GetOrder

type Order struct {
	Id                string         `db:"id" form:"id" json:"id"`
	Order_number      string         `db:"order_number" form:"order_number" json:"order_number"`
	User_id           string         `db:"user_id" form:"user_id" json:"user_id"`
	Payment_method_id int            `db:"payment_method_id" form:"payment_method_id" json:"payment_method_id"`
	Movie_time_id     int            `db:"movie_time_id" form:"movie_time_id" json:"movie_time_id"`
	Seat_count        int            `db:"seat_count" form:"seat_count" json:"seat_count"`
	Ticket_status     string         `db:"ticket_status" form:"ticket_status" json:"ticket_status"`
	Orders            []OrderDetails `json:"orders"`
	Created_at        *time.Time     `db:"created_at" form:"created_at" json:"created_at"`
	Updated_at        *time.Time     `db:"updated_at" form:"updated_at" json:"updated_at"`
}
