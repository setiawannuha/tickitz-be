package models

var schemaOrderDetails = `
CREATE TABLE order_details (
    id SERIAL PRIMARY KEY,
    order_id UUID,
    seat_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (seat_id) REFERENCES seats(id)
);
`