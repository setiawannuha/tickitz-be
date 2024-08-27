package models

var schemaOrderDetails = `
CREATE TABLE public.order_details (
	id serial4 NOT NULL,
	order_id uuid NULL,
	seat_id int4 NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT order_details_pkey PRIMARY KEY (id),
	CONSTRAINT order_details_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id),
	CONSTRAINT order_details_seat_id_fkey FOREIGN KEY (seat_id) REFERENCES public.seats(id)
);
`

type OrderDetails struct {
	Order_id string `db:"order_id" form:"order_id" json:"order_id"`
	Seat_id  int    `db:"seat_id" form:"seat_id" json:"seat_id"`
}

type GetOrderDetails struct {
	Order_id *string `db:"order_id" form:"order_id" json:"order_id"`
	Seat_id  *int    `db:"seat_id" form:"seat_id" json:"seat_id"`
}
