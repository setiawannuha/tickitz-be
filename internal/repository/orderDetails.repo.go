package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderDetailsRepositoryInterface interface {
	CreateOrderDetails(order_id string, orders []models.OrderDetails)(string, error)
	GetDetailOrder(order_id string)(*[]models.GetOrderDetails, error)
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

func (r *OrderDetailsRepository) GetDetailOrder(order_id string)(*[]models.GetOrderDetails, error){
	query := `select order_id, seat_id from order_details where order_id=$1`
	data := []models.GetOrderDetails{}

	err := r.Select(&data, query, order_id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

