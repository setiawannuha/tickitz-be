package routers

import (
	"khalifgfrz/coffee-shop-be-go/internal/handlers"
	"khalifgfrz/coffee-shop-be-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func orderRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/orders")

	var orderRepo repository.OrderRepositoryInterface = repository.NewOrderRepository(d)
	var orderDetailsRepo repository.OrderDetailsRepositoryInterface = repository.NewOrderDetailsRepository(d)
	handler := handlers.NewOrderHandler(orderRepo, orderDetailsRepo)

	router.POST("/", handler.CreateOrder)
	router.GET("/history", handler.FetchHistory)
	router.GET("/", handler.FetchAll)
	router.GET("/:id", handler.FetchDetail)
}