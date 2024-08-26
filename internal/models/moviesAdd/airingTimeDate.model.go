package moviesAdd

import "time"

var schemaAiringTimeDate = `
CREATE TABLE public.airing_time_date (
    id SERIAL PRIMARY KEY,
    airing_time_id INT,
    date_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (airing_time_id) REFERENCES airing_time(id),
    FOREIGN KEY (date_id) REFERENCES airing_date(id)
);
`

type AiringTimeDate struct {
	Id             int       `db:"id" json:"id"`
	Airing_time_id int       `db:"airing_time_id" json:"airing_time_id"`
	Date_id        int       `db:"date_id" json:"date_id"`
	Created_at     time.Time `db:"created_at" json:"created_at"`
	Updated_at     time.Time `db:"updated_at" json:"updated_at"`
}
