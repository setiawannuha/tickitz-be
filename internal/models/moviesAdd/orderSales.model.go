package moviesAdd

import (
	"time"
)

type GetSales struct {
	Date       string     `db:"date" json:"date"`
	Sales      int        `db:"sales" json:"sales"`
	Created_at *time.Time `db:"created_at" json:"created_at,omitempty"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
