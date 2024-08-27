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
	Order_number      *string           `db:"order_number" json:"order_number" form:"order_number"`
	User_id           *string           `db:"user_id" json:"user_id,omitempty" form:"user_id"`
	Payment_method_id *int              `db:"payment_method_id" json:"payment_method_id,omitempty" form:"payment_method_id"`
	Movie_id          *string           `db:"movie_id" json:"movie_id,omitempty" form:"movie_id"`
	Movie_title       *string           `db:"movie_title" json:"movie_title,omitempty" form:"movie_title"`
	Date              *string           `db:"date" json:"date" form:"date"`
	Time              *string           `db:"time" json:"time" form:"time"`
	Seat_count        *int              `db:"seat_count" json:"seat_count" form:"seat_count"`
	Ticket_status     *string           `db:"ticket_status" json:"ticket_status" form:"ticket_status"`
	Genres            []byte            `db:"genres" json:"genres"` // Changed to a slice of strings
	Total             *int              `db:"total" json:"total" form:"total"`
	Orders            []GetOrderDetails `json:"orders"`
	Created_at        *time.Time        `db:"created_at" json:"created_at,omitempty" form:"created_at"`
	Updated_at        *time.Time        `db:"updated_at" json:"updated_at,omitempty" form:"updated_at"`
}

type GetOrders []GetOrder

// Order represents the order data with additional details.
type Order struct {
	Id                string         `db:"id" json:"id" form:"id"`
	Order_number      string         `db:"order_number" json:"order_number" form:"order_number"`
	User_id           string         `db:"user_id" json:"user_id" form:"user_id"`
	Payment_method_id int            `db:"payment_method_id" json:"payment_method_id" form:"payment_method_id"`
	Movie_id          *string        `db:"movie_id" json:"movie_id" form:"movie_id"`
	Date              *string        `db:"date" json:"date" form:"date"`
	Time              *string        `db:"time" json:"time" form:"time"`
	Seat_count        int            `db:"seat_count" json:"seat_count" form:"seat_count"`
	Ticket_status     string         `db:"ticket_status" json:"ticket_status" form:"ticket_status"`
	Total             int            `db:"total" json:"total" form:"total"`
	Orders            []OrderDetails `json:"orders"` // Assumes OrderDetails is already defined
	Created_at        *time.Time     `db:"created_at" json:"created_at" form:"created_at"`
	Updated_at        *time.Time     `db:"updated_at" json:"updated_at" form:"updated_at"`
}
