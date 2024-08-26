package repository

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepositoryInterface interface {
	CreateData(body *models.Order)(string, error)
	GetAllData()(*models.GetOrders, error)
	GetDetailData(id string)(*models.GetOrder, error)
	GetHistoryOrder(id string)(*models.GetOrder, error)
}

type OrderRepository struct {
	*sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) CreateData(body *models.Order)(string, error){
	query := `INSERT INTO public.orders
(order_number, user_id, payment_method_id, movie_time_id, seat_count, ticket_status)
VALUES($1, $2, $3, $4, $5, $6) returning id;`
	params := []interface{}{
		body.Order_number,
		body.User_id,
		body.Payment_method_id,
		body.Movie_time_id,
		body.Seat_count,
		body.Ticket_status,
	}

	var id string
	err := r.Get(&id, query, params...)
	if err != nil {
		return "", err
	}
    return id, nil
}

func (r *OrderRepository) GetAllData()(*models.GetOrders, error){
	query := `select id, order_number, user_id, payment_method_id, movie_time_id, seat_count, ticket_status from order`
	data := models.GetOrders{}

	err := r.Select(&data, query)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *OrderRepository) GetDetailData(id string)(*models.GetOrder, error){
	query := `select id, order_number, user_id, payment_method_id, movie_time_id, seat_count, ticket_status from order where id=$1`
	data := models.GetOrder{}

	err := r.Get(&data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *OrderRepository) GetHistoryOrder(id string)(*models.GetOrder, error){
	query := `select o.id, o.order_number, o.user_id, o.payment_method_id, o.movie_time_id, o.seat_count, o.ticket_status from order o join public.users u on o.user_id = u.id where u.id=$1`
	data := models.GetOrder{}

	err := r.Get(&data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}


