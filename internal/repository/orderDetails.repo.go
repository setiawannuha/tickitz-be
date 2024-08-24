package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderDetailsRepositoryInterface interface {
	CreateOrderDetails(order_id string, orders []models.OrderDetails)(string, error)
}

type OrderDetailsRepository struct {
	*sqlx.DB
}

func NewOrderDetailsRepository(db *sqlx.DB) *OrderDetailsRepository {
	return &OrderDetailsRepository{db}
}

func (r *OrderDetailsRepository) CreateOrderDetails(order_id string, orders []models.OrderDetails)(string, error){
	query := `INSERT INTO public.order_details
	(order_id, seat_id)
	VALUES(:order_id, :seat_id);`

	for _, order:=range orders {
		data := map[string]interface{}{
			"order_id":order_id,
			"seat_id":order.Seat_id,
		}
		_, err := r.NamedExec(query, data)
		if err != nil {
			return "", err
		}
	}

	
    return "Order created", nil
}