package handlers

import (
	"khalifgfrz/coffee-shop-be-go/internal/models"
	"khalifgfrz/coffee-shop-be-go/internal/repository"
	"khalifgfrz/coffee-shop-be-go/pkg"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	repository.OrderRepositoryInterface
	repository.OrderDetailsRepositoryInterface
	repository.PaymentsRepoInterface
}

func NewOrderHandler(orderRepo repository.OrderRepositoryInterface, orderDetailsRepo repository.OrderDetailsRepositoryInterface, payments repository.PaymentsRepoInterface) *OrderHandler {
	return &OrderHandler{orderRepo, orderDetailsRepo, payments}
}

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.Order{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("Create order failed 1", "Error")
		return
	}

	orderID, err := h.CreateData(&body)
	if err != nil {
		response.BadRequest("Create order failed 2", "Error")
		return
	}

	result, err := h.CreateOrderDetails(orderID, body.Orders)
	if err != nil {
		response.BadRequest("Create order failed 3", "Failed to create order details")
		return
	}

	response.Created("Create order success", result)
}

func (h *OrderHandler) FetchAll(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	orders, err := h.GetAllData()
	if err != nil {
		response.InternalServerError("Get data failed", "Error")
		return
	}
	for i := range *orders {
		order := &(*orders)[i]

		orderDetails, err := h.GetDetailOrder(order.Id)
		if err != nil {
			response.InternalServerError("Get data failed", "Error")
			return
		}
		order.Orders = orderDetails
	}
	response.Success("Get data success", orders)
}

func (h *OrderHandler) FetchDetail(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	order, err := h.GetDetailData(id)

	if err != nil {
		response.InternalServerError("Get data failed", "Error")
		return
	}

	orderID := order.Id

	orderDetails, err := h.GetDetailOrder(orderID)

	if err != nil {
		response.InternalServerError("Get data failed", "Error")
		return
	}

	order.Orders = orderDetails

	response.Success("Get data success", order)
}

func (h *OrderHandler) FetchHistory(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	userID, exists := ctx.Get("id")

	if !exists {
		response.InternalServerError("User id not found", nil)
		return
	}

	id := userID.(string)

	history, err := h.GetHistoryOrder(id)
	if err != nil {
		response.InternalServerError("Get data failed", "Error")
		return
	}

	// Iterate through each order and get its details
	for i, order := range history {
		orderID := order.Id

		orderDetails, err := h.GetDetailOrder(orderID)
		if err != nil {
			response.InternalServerError("Get data failed", "Error")
			return
		}

		// Populate the Orders field with the details
		history[i].Orders = orderDetails
	}

	response.Success("Get data success", history)
}

//additional

func (h *OrderHandler) GetPayments(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	data, err := h.GetAllPayments()
	if err != nil {
		response.InternalServerError("Internal Server Error", err.Error())
	}
	if len(data) == 0 {
		response.NotFound("Data Not Found", "No datas available for the given criteria")
		return
	}

	response.Success("Data retrieved successfully", data)
}
