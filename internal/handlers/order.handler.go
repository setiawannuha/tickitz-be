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
}

func NewOrderHandler(orderRepo repository.OrderRepositoryInterface , orderDetailsRepo repository.OrderDetailsRepositoryInterface) *OrderHandler {
	return &OrderHandler{orderRepo , orderDetailsRepo}
}

func (h *OrderHandler) CreateOrder (ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.Order{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("Create order failed", err.Error())
		return
	}

	orderID, err := h.CreateData(&body)
	if err != nil {
		response.BadRequest("Create order failed", err.Error())
		return
	}
	
	result, err := h.CreateOrderDetails(orderID, body.Orders)
	if err != nil {
		response.BadRequest("Create order failed", err.Error())
		return
	}

	response.Created("Create order success", result)
}