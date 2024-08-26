package models

var schemaOrderDetails = `
CREATE TABLE order_details (
    id SERIAL PRIMARY KEY,
    order_id UUID,
    seat_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES ord   ers(id),
    FOREIGN KEY (seat_id) REFERENCES seats(id)
);
`

type OrderDetails struct {
	Order_id string `db:"order_id" form:"order_id" json:"order_id"`
	Seat_id  int    `db:"seat_id" form:"seat_id" json:"seat_id"`
}
